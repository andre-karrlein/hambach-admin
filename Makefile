build:
	@go build -o hambach-admin ./app

wasm:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm ./app

run: build wasm
	export GOOGLE_APPLICATION_CREDENTIALS="./hambach.json" && ./hambach-admin