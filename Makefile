all: check install docker

setup:
	@echo "setup"
	@go get github.com/golang/lint/golint
	@go get golang.org/x/tools/cmd/goimports
	@go get github.com/golang/dep/cmd/dep
	dep ensure

check:
	@echo "check"
	@ginkgo rota_test

install:
	@echo "install"
	@mkdir -p ./bin
	@env GOOS=linux GOARCH=amd64 go install -v ./next-to-support
	@cp ~/go/bin/linux_amd64/next-to-support ./bin
	

docker:
	@echo "docker"
	@docker build -t support-bot:5 .
