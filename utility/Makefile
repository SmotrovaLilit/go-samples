build:
	go build -o fixtext  -v ./cmd/fixtext/main.go

run:
	ollama pull mistral-nemo
	AI_TYPE=ollama go run ./cmd/fixtext/main.go