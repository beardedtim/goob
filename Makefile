BINARY_NAME_PREFIX=goob
HTTP_BINARY_PREFIX=http

test:
	go test ./...

dev-http:
	go run cmd/http/main.go

tidy:
	go mod tidy
	go fmt ./...

clean:
	rm -rf ./dist
	mkdir ./dist
	touch ./dist/.gitkeep

build-http:
	GOARCH=amd64 GOOS=darwin go build -o ./dist/${BINARY_NAME_PREFIX}-${HTTP_BINARY_PREFIX}-darwin cmd/http/main.go
 	GOARCH=amd64 GOOS=linux go build -o ./dist/${BINARY_NAME_PREFIX}-${HTTP_BINARY_PREFIX}-linux cmd/http/main.go
 	GOARCH=amd64 GOOS=windows go build -o ./dist/${BINARY_NAME_PREFIX}-${HTTP_BINARY_PREFIX}-windows cmd/http/main.go

build: clean build-http