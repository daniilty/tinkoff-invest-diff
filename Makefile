build:
	go build github.com/daniilty/tinkoff-invest-diff/cmd/diff
build_docker:
	docker build -f docker/Dockerfile .
