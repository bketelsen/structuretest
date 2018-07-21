.PHONY: wasm
wasm:
	GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o ./public/app.wasm ./client
	rsync -a --prune-empty-dirs --include '*/' --include '*.html' --exclude '*' components/ templates/

.PHONY: server-app
server-app:
	GOOS=darwin GOARCH=amd64 ~/gowasm/bin/go build -o server-app ./server

.PHONY: run
run: wasm server-app
	./server-app
