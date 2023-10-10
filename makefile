build:
	env GOOS=js GOARCH=wasm go build -o views/game.wasm ./cmd/main.go