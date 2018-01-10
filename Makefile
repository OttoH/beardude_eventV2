init: deps
dev: dev
db: mongo
build: build

.PHONY: deps
deps:
	brew install glide #curl https://glide.sh/get | sh
	glide install

.PHONY: dev
dev:
	go run main.go routers.go authMiddleware.go

.PHONY: build
build:
	rm -rf ./bin/beardudev2
	go build -o ./bin/beardudev2 main.go routers.go authMiddleware.go
	GIN_MODE=release
	./bin/beardudev2

.PHONY: mongo
mongo:
		mongod --dbpath data/db
