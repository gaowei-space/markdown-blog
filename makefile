GO111MODULE=on

.PHONY: build
build: bindata markdown-blog

.PHONY: docker-push
docker-push: package-all
	docker buildx build --platform linux/arm64,linux/amd64 -t willgao/markdown-blog:latest . --push

.PHONY: docker-build
docker-build: package-all
	docker build -t willgao/markdown-blog:dev -f ./Dockerfile.Develop .

.PHONY: bindata
bindata:
	go install github.com/go-bindata/go-bindata/v3/go-bindata@latest
	go generate ./...

.PHONY: markdown-blog
markdown-blog:
	go build $(RACE) -o bin/markdown-blog ./

.PHONY: build-race
build-race: enable-race build

.PHONY: run
run: build
	./bin/markdown-blog web --config ./config/config.yml

.PHONY: run-race
run-race: enable-race run

.PHONY: test
test:
	go test $(RACE) ./...

.PHONY: test-race
test-race: enable-race test

.PHONY: enable-race
enable-race:
	$(eval RACE = -race)

.PHONY: package
package: build
	bash ./package.sh

.PHONY: package-all
package-all: build
	bash ./package.sh -p 'linux darwin windows' -a 'amd64 arm64'

.PHONY: clean
clean:
	rm bin/markdown-bolg
