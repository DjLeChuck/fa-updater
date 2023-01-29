# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet all code
.PHONY: audit
audit: tidy
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...

## vendor: tidy dependencies
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/cmd: build the cmd application
.PHONY: build/cmd
build/cmd:
	@echo 'Building cmd...'
	GOOS=linux GOARCH=amd64 go build -ldflags="-s" -o=./bin/linux_amd64/fa-updater ./
	GOOS=windows GOARCH=amd64 go build -ldflags="-s" -o=./bin/windows_amd64/fa-updater.exe ./
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s" -o=./bin/darwin_amd64/fa-updater ./
