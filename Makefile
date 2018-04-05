.PHONY: all
all: clean build build-docker push

clean:
	go clean
build:
	GOOS=linux go build -v

build-docker:
	docker build -t gauravthadani/istiosample .

push:
	docker push gauravthadani/istiosample
