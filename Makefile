server:
	go run .
build-wasm b:
	GOOS=js GOARCH=wasm go build -o ./out/main.wasm ./app
get-wasm-js:
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .