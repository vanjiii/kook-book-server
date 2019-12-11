COVER ?= coverage.html

all: clean dep migrate build

migrate: ## Executes database migration.
	goose -dir ./pkg/migrations postgres "dbname=kookbook sslmode=disable" up

build: ## Build binary file.
	go install ./...

dep: ## Install dependencies.
	go get ./...
	go mod tidy
	which goose > /dev/null || go get github.com/pressly/goose/cmd/goose@v2.6.0
	which staticcheck > /dev/null || go get honnef.co/go/tools/cmd/staticcheck@2019.2.3

test:
	go test -coverprofile=coverage.out -covermode=count ./...
	go vet ./...
	staticcheck ./...

cover: test ## Generates and open HTML code coverage report.
	go tool cover -html=coverage.out -o $(COVER)
	xdg-open $(COVER)

clean:
	rm -f coverage.out
	rm -f $(COVER)

help: ## Display help screen.
	@grep -h -E '^[\.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
