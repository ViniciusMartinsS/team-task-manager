migrate:

setup:
	# @go mod download
	@go run cmd/database/main.go
	# @go get -d gotest.tools/gotestsum
	# @go get -d github.com/go-courier/husky/cmd/husky && husky init

migrate:
	@go run cmd/database/main.go

run:
	@go run ./main.go

lint:
	@staticcheck ./...
