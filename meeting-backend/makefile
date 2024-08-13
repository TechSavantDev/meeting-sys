BINARY_FILE = backend
ENTRY_FILE = cmd/meeting-system/main.go

all:
	@make build
	@echo "Running $(BINARY_FILE)..."
	@./$(BINARY_FILE)

build:
	@echo "Building $(BINARY_FILE)..."
	@go mod tidy
	@go build -o $(BINARY_FILE) $(ENTRY_FILE)
	@echo "Done."

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_FILE)
	@echo "Done."
