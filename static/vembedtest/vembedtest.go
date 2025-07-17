// Ollama Embedding Determinism Test
//
// This program tests whether Ollama embedding models produce deterministic results
// by generating embeddings multiple times for the same input and comparing them.
//
// Usage: go run main.go [flags]

package main

import (
	"bytes"
	"encoding/json"
	"flag"
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

// Configuration flags
var (
	ollamaURL  = flag.String("url", "http://localhost:11434/api/embed", "Full Ollama embedding API endpoint URL (include /api/embed)")
	iterations = flag.Int("iterations", 5, "Number of times to test each embedding for determinism")
	verbose    = flag.Bool("verbose", false, "Enable verbose debug output")
	help       = flag.Bool("help", false, "Show help message")
)

var (
	// Test models - update these to match your available models
	// Run "ollama list" to see your available models
	testModels = []string{
		"all-minilm:latest",
		"bge-large:latest",
		"bge-m3:latest",
		"granite-embedding:latest",
		"mxbai-embed-large:latest",
		"nomic-embed-text:latest",
		"paraphrase-multilingual:latest",
		"snowflake-arctic-embed2:latest",
		"snowflake-arctic-embed:latest",
		// Add your models here, e.g.:
		// "all-minilm:latest",
		// "bge-large:latest",
	}

	// Test strings with varying lengths and complexity
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
	flag.Parse()

	if *help {
		fmt.Println("🧪 Ollama Embedding Determinism Test")
		fmt.Println(strings.Repeat("=", 50))
		fmt.Println()
		fmt.Println("This tool tests whether Ollama embedding models produce deterministic results")
		fmt.Println("by generating embeddings multiple times for the same input and comparing them.")
		fmt.Println()
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  # Run with defaults (localhost)")
		fmt.Println("  go run main.go")
		fmt.Println()
		fmt.Println("  # Run 10 iterations with verbose output")
		fmt.Println("  go run main.go -iterations 10 -verbose")
		fmt.Println()
		fmt.Println("  # Connect to remote Ollama instance")
		fmt.Println("  go run main.go -url http://zima:11434/api/embed")
		fmt.Println()
		fmt.Println("Setup:")
		fmt.Println("  1. Make sure Ollama is running: ollama serve")
		fmt.Println("  2. Check available models: ollama list")
		fmt.Println("  3. Update testModels in code to match your installed models")
		fmt.Println("  4. Test manually: curl http://localhost:11434/api/embed -d '{\"model\":\"your-model\",\"input\":\"test\"}'")
		return
	}

	fmt.Println("🧪 Testing Ollama Embedding Determinism")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Testing %d models with %d strings, %d iterations each\n", len(testModels), len(testStrings), *iterations)
	fmt.Printf("Ollama URL: %s\n", *ollamaURL)
	if *verbose {
		fmt.Printf("Verbose mode: enabled\n")
	}
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

			result := testEmbeddingDeterminism(model, input, *iterations)
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

	if *verbose {
		log.Printf("Request to %s: %s", *ollamaURL, string(jsonData))
	}

	resp, err := http.Post(*ollamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if *verbose {
		log.Printf("Response status: %d, body: %s", resp.StatusCode, string(body))
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
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

	// Calculate actual maximum widths needed
	maxModelLen := len("MODEL")
	maxInputLen := len("INPUT TEXT")

	for _, result := range results {
		if len(result.Model) > maxModelLen {
			maxModelLen = len(result.Model)
		}
		// Truncate input text for display and measure actual display length
		inputDisplay := result.Input
		if len(inputDisplay) > 35 {
			inputDisplay = inputDisplay[:32] + "..."
		}
		if len(inputDisplay) > maxInputLen {
			maxInputLen = len(inputDisplay)
		}
	}

	// Ensure minimum column widths and add padding
	if maxModelLen < 10 {
		maxModelLen = 10
	}
	if maxInputLen < 15 {
		maxInputLen = 15
	}

	// Fixed widths for remaining columns
	const deterministicWidth = 13
	const embedSizeWidth = 10
	const iterationsWidth = 10

	// Create format strings
	headerFormat := fmt.Sprintf("| %%-%ds | %%-%ds | %%-%ds | %%%ds | %%%ds |\n",
		maxModelLen, maxInputLen, deterministicWidth, embedSizeWidth, iterationsWidth)
	rowFormat := fmt.Sprintf("| %%-%ds | %%-%ds | %%-%ds | %%%dd | %%%dd |\n",
		maxModelLen, maxInputLen, deterministicWidth, embedSizeWidth, iterationsWidth)

	// Calculate total width for borders
	totalWidth := maxModelLen + maxInputLen + deterministicWidth + embedSizeWidth + iterationsWidth + 16 // +16 for separators and padding

	// Top border
	fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")

	// Header
	fmt.Printf(headerFormat, "MODEL", "INPUT TEXT", "DETERMINISTIC", "EMBED SIZE", "ITERATIONS")

	// Header separator
	fmt.Printf("|%s+%s+%s+%s+%s|\n",
		strings.Repeat("-", maxModelLen+2),
		strings.Repeat("-", maxInputLen+2),
		strings.Repeat("-", deterministicWidth+2),
		strings.Repeat("-", embedSizeWidth+2),
		strings.Repeat("-", iterationsWidth+2))

	// Data rows
	for _, result := range results {
		inputDisplay := result.Input
		if len(inputDisplay) > 35 {
			inputDisplay = inputDisplay[:32] + "..."
		}

		deterministic := "NO"
		if result.IsDeterministic {
			deterministic = "YES"
		}

		fmt.Printf(rowFormat,
			result.Model,
			inputDisplay,
			deterministic,
			result.EmbeddingSize,
			result.Iterations)
	}

	// Bottom border
	fmt.Println("+" + strings.Repeat("-", totalWidth-2) + "+")
	fmt.Println()

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
