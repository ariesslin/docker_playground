export tag=v2.0
root:
	export ROOT=docker_playground/homework

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t ariesslin/httpserver_test:${tag} .

push: release
	echo "pushing ariesslin/httpserver_test"
	docker push ariesslin/httpserver_test:${tag}
