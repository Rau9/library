# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

all: test package

# Build manager binary
build: fmt vet
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o main main.go

# Run tests
test: fmt vet docker-compose-up
	go test ./...

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

tidy:
	go mod tidy

run: docker-compose-up
	go run main.go

docker-compose-up:
	docker-compose up -d
	sleep 5

docker-compose-down:
	docker-compose down

docker-build: fmt vet
	docker build -t library:local .

clean: docker-compose-down
	rm -f main
