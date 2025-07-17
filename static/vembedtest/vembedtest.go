package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// EmbedRequest represents the request structure for Ollama's embed API
type EmbedRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

// EmbedResponse represents the response structure from Ollama's embed API
type EmbedResponse struct {
	Embeddings [][]float64 `json:"embeddings"`
}

// TestResult holds the results of determinism testing for a model-input combination
type TestResult struct {
	Model           string
	Input           string
	IsDeterministic bool
	Iterations      int
	EmbeddingSize   int
}

// Configuration
const (
	ollamaURL  = "http://localhost:11434/api/embed"
	iterations = 5     // Number of times to test each embedding
	verbose    = false // Set to true for debug output
)

var (
	// Test models - adjust based on what you have available
	// Run "ollama list" to see your available models
	testModels = []string{
		"mxbai-embed-large",
		"nomic-embed-text",
		"all-minilm",
		// Add your models here, e.g.:
		// "all-minilm:latest",
		// "bge-large:latest",
	}

	// Test strings
	testStrings = []string{
		"Hello, world!",
		"The quick brown fox jumps over the lazy dog.",
		"Artificial intelligence is transforming the world.",
		"Vector embeddings represent semantic meaning of text.",
		"Short text",
		"This is a longer sentence with more complex vocabulary and structure to test how embedding models handle varying text lengths and complexity.",
	}
)

func main() {
	fmt.Println("🧪 Testing Ollama Embedding Determinism")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Testing %d models with %d strings, %d iterations each\n", len(testModels), len(testStrings), iterations)
	fmt.Printf("💡 Tip: Run 'ollama list' to see available models and update testModels in the code\n")
	fmt.Println()

	var allResults []TestResult

	for _, model := range testModels {
		fmt.Printf("Testing model: %s\n", model)

		for i, input := range testStrings {
			fmt.Printf("  [%d/%d] Testing string: \"%.50s%s\"\n",
				i+1, len(testStrings), input, func() string {
					if len(input) > 50 {
						return "..."
					}
					return ""
				}())

			result := testEmbeddingDeterminism(model, input, iterations)
			allResults = append(allResults, result)

			status := "✅ DETERMINISTIC"
			if !result.IsDeterministic {
				status = "❌ NON-DETERMINISTIC"
			}
			fmt.Printf("    Result: %s (embedding size: %d)\n", status, result.EmbeddingSize)
		}
		fmt.Println()
	}

	// Display summary table
	displayResultsTable(allResults)
}

func testEmbeddingDeterminism(model, input string, iterations int) TestResult {
	var embeddings [][]float64

	for i := 0; i < iterations; i++ {
		embedding, err := getEmbedding(model, input)
		if err != nil {
			log.Printf("Error getting embedding for model %s: %v", model, err)
			return TestResult{
				Model:           model,
				Input:           input,
				IsDeterministic: false,
				Iterations:      i,
				EmbeddingSize:   0,
			}
		}
		embeddings = append(embeddings, embedding)

		// Small delay to avoid overwhelming the server
		time.Sleep(100 * time.Millisecond)
	}

	// Check if all embeddings are identical
	isDeterministic := true
	if len(embeddings) > 1 {
		firstEmbedding := embeddings[0]
		for i := 1; i < len(embeddings); i++ {
			if !reflect.DeepEqual(firstEmbedding, embeddings[i]) {
				isDeterministic = false
				break
			}
		}
	}

	embeddingSize := 0
	if len(embeddings) > 0 {
		embeddingSize = len(embeddings[0])
	}

	return TestResult{
		Model:           model,
		Input:           input,
		IsDeterministic: isDeterministic,
		Iterations:      iterations,
		EmbeddingSize:   embeddingSize,
	}
}

func getEmbedding(model, input string) ([]float64, error) {
	reqBody := EmbedRequest{
		Model: model,
		Input: input,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		if verbose {
			log.Printf("API Error Response: %s", string(body))
		}
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if verbose {
		log.Printf("Raw response for model %s: %s", model, string(body))
	}

	var embedResp EmbedResponse
	if err := json.Unmarshal(body, &embedResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(embedResp.Embeddings) == 0 {
		return nil, fmt.Errorf("no embeddings returned")
	}

	return embedResp.Embeddings[0], nil
}

func displayResultsTable(results []TestResult) {
	fmt.Println("📊 SUMMARY RESULTS")
	fmt.Println(strings.Repeat("=", 80))

	// Calculate column widths
	maxModelLen := len("MODEL")
	maxInputLen := len("INPUT TEXT")
	for _, result := range results {
		if len(result.Model) > maxModelLen {
			maxModelLen = len(result.Model)
		}
		if len(result.Input) > maxInputLen && len(result.Input) <= 40 {
			maxInputLen = len(result.Input)
		}
	}
	if maxInputLen > 40 {
		maxInputLen = 40
	}

	// Header
	fmt.Printf("%-*s | %-*s | %-13s | %-10s | %s\n",
		maxModelLen, "MODEL",
		maxInputLen, "INPUT TEXT",
		"DETERMINISTIC",
		"EMBED SIZE",
		"ITERATIONS")
	fmt.Println(strings.Repeat("-", maxModelLen) + "-+-" +
		strings.Repeat("-", maxInputLen) + "-+-" +
		"-------------+-" +
		"----------+-" +
		"----------")

	// Results
	for _, result := range results {
		inputDisplay := result.Input
		if len(inputDisplay) > 40 {
			inputDisplay = inputDisplay[:37] + "..."
		}

		deterministic := "❌ NO"
		if result.IsDeterministic {
			deterministic = "✅ YES"
		}

		fmt.Printf("%-*s | %-*s | %-13s | %-10d | %d\n",
			maxModelLen, result.Model,
			maxInputLen, inputDisplay,
			deterministic,
			result.EmbeddingSize,
			result.Iterations)
	}

	fmt.Println()

	// Statistics
	totalTests := len(results)
	deterministicCount := 0
	modelStats := make(map[string]int)
	modelTotal := make(map[string]int)

	for _, result := range results {
		if result.IsDeterministic {
			deterministicCount++
			modelStats[result.Model]++
		}
		modelTotal[result.Model]++
	}

	fmt.Printf("📈 STATISTICS:\n")
	fmt.Printf("Total tests: %d\n", totalTests)
	fmt.Printf("Deterministic: %d (%.1f%%)\n", deterministicCount, float64(deterministicCount)/float64(totalTests)*100)
	fmt.Printf("Non-deterministic: %d (%.1f%%)\n", totalTests-deterministicCount, float64(totalTests-deterministicCount)/float64(totalTests)*100)
	fmt.Println()

	fmt.Printf("📋 BY MODEL:\n")
	for model, total := range modelTotal {
		deterministic := modelStats[model]
		fmt.Printf("  %s: %d/%d deterministic (%.1f%%)\n",
			model, deterministic, total, float64(deterministic)/float64(total)*100)
	}
}
