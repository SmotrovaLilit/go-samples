## Install to use with ollama

- Install ollama [https://ollama.com](https://ollama.com)

- Install the mistral-nemo model

```bash
ollama pull mistral-nemo
```
## Run with ollama

```bash
AI_TYPE=ollama go run ./cmd/fixtext/main.go
```

## Run with openai

```bash
AI_TYPE=openai OPEN_AI_API_KEY=YOUR_KEY go run ./cmd/fixtext/main.go
```

## Additional information
### Ollama API

[ollama api](https://github.com/ollama/ollama/blob/main/docs/api.md#generate-a-completion)

### Models

[library](https://ollama.com/library)

### Run mistral-nemo

```bach
ollama run mistral-nemo
ollama ps
```

### Check ollama logs

```bash
tail -f ~/.ollama/logs/server.log
```