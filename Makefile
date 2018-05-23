BUILD_VERBOSE := -v

TEST_VERBOSE := -v

DOCKER_IMAGE_NAME := johnbelamaric/grpc-whoamid
DOCKER_VERSION := latest

all: client server

# Phony this to ensure we always build the binary.
# TODO: Add .go file dependencies.
.PHONY: client
client: deps
	cd grpc-whoami && go build $(BUILD_VERBOSE) -ldflags="-s -w"

.PHONY: server
server: deps
	cd grpc-whoamid && go build $(BUILD_VERBOSE) -ldflags="-s -w"

.PHONY: docker
docker: deps
	cd grpc-whoamid && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w"
	cd grpc-whoamid && docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_VERSION) .

.PHONY: deps
deps:
	cd grpc-whoamid && go get ${BUILD_VERBOSE}
	cd grpc-whoami && go get ${BUILD_VERBOSE}

.PHONY: pb
pb:
	protoc --go_out=plugins=grpc:pb whoami.proto

.PHONY: clean
clean:
	go clean
	rm -f grpc-whoami/grpc-whoami grpc-whoamid/grpc-whoamid
