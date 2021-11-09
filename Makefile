migrate:

setup:
	@go mod download
	@go run cmd/database/main.go
	@go get -d gotest.tools/gotestsum
	@go get -d github.com/go-courier/husky/cmd/husky && husky init

docker:
	@docker-compose up --build

migrate:
	@go run cmd/database/main.go

run:
	@go run ./main.go

lint:
	@staticcheck ./...

tests:
	@gotestsum -f pkgname ./test/application/...

coverage:
	@go test -v -cover -coverprofile=r.out -coverpkg ./internal/... ./test/...
	@go tool cover -html=r.out -o report.html
	@rm -f r.out
	@google-chrome --new-window report.html
