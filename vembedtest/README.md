# quick embedding determinism test

Q: Are vector embeddings deterministic?

Discussion:

* [So text embeddings are kind of random like LLMs?](https://news.ycombinator.com/item?id=39958719)
* [Why `OpenAI Embedding` return different vectors for the same text input?](https://community.openai.com/t/why-openai-embedding-return-different-vectors-for-the-same-text-input/144495/4)
* ...

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

![](vembedtest.gif)
