# quick embedding determinism test

Q: Are vector embeddings deterministic?

We do a quick test using ollama.

```
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
}
```
