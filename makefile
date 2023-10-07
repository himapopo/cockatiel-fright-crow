build:
	env GOOS=js GOARCH=wasm go build -o game.wasm ./cmd/main.go