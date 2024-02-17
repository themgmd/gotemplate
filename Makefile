.PHONY: install_tooling generate migration cover convey

install_tooling:
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/smartystreets/goconvey@latest

generate:
	go generate ./...

migration:
	@read -p "Enter Migration Name:" migration_name; \
	sql-migrate new --config=dbconfig.yml -env=gotemplate $$migration_name

cover:
	go test -short -count=1 -race -coverprofile coverage.out ./internal/...
	go tool cover -html coverage.out -o cover.html
	rm coverage.out

convey:
	goconvey -excludedDirs=vendor -port 8090