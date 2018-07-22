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

db-up:
	docker run --name=structuretest_db -d -p 5432:5432 -e POSTGRES_DB=structuretest_development postgres
	sleep 10
	buffalo db migrate up
	docker ps | grep structuretest_db 

db-down:
	docker stop structuretest_db
	docker rm structuretest_db
