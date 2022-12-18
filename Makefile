.PHONY: tidy docker docker-release build-release

tidy:
	go mod tidy

docker:
	docker build . -t megrez

docker-release: docker
	docker push megrez@latest

build-release: tidy docker-release
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/megrez-linux-amd64 main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/megrez-linux-arm64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/megrez-darwin-amd64 main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/megrez-darwin-arm64 main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/megrez-windows-amd64.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/megrez-windows-arm64.exe main.go
