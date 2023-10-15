all: vet lint fmt race test build	

dependencies:
	@echo "Executing '$@'..."
	@go get -v -d ./...

vet: dependencies
	@echo "Executing '$@'..."
	@go vet ./...

# Run linter (staticcheck)
lint: dependencies
	@echo "Executing '$@'..."
	@staticcheck ./...

# Format the files
fmt: dependencies
	@echo "Executing '$@'..."
	@go fmt ./...

# Run data race detector
race: dependencies
	@echo "Executing '$@'..."
	@go test -v -race -short ./...

test: dependencies
	@echo "Executing '$@'..."
	@go test -v -short -count=1 ./...

build: dependencies
	@echo "Executing '$@'..."
	@go build -o ./bin/bridge-validator cmd/main.go

clean:
	@echo "Executing '$@'..."
	@rm -rf ./bin/