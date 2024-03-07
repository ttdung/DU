build: build-app
run:
	make build-app && ./du --config config.json
build-app:
	go build -o du
clean:
	rm -f du && go clean --modcache && go mod tidy
