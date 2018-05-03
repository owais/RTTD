.PHONY: build install test run lint vet restart kill watch

pid = /tmp/RTTD.pid
run_cmd = ./build/timezones

VERSION ?= $(shell git tag | tail -1)

LEVEL ?= development

ld_flags = -ldflags="-X 'main.version=$(version)'"

all: install build

install:
	go get github.com/golang/dep/cmd/dep
	go get github.com/phogolabs/parcello
	go install github.com/phogolabs/parcello/cmd/parcello
	@echo "\n======\nNote: https://github.com/emcrisostomo/fswatch is required to use 'make watch'\n======\n"
	@echo "Done! Run 'make watch' or 'make run'"

build: clean
	@echo "building..."
	@gopherjs build -m -o static/main.js ./cmd/frontend/
	@go generate ./...
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/rttd_linux $(ld_flags) ./cmd/server
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/rttd_darwin $(ld_flags) ./cmd/server

clean:
	@rm -rf build/

lint:
	@echo "linting..."
	@go list ./... | grep -v /vendor/ | xargs -L1 golint --min_confidence 0

vet:
	@echo "vetting..."
	@go list ./... | grep -v /vendor/ | xargs -L1 go vet

test: lint vet
	@echo "testing..."
	@go test

run:
	@cd build; $(run_cmd)

watch: restart
	@fswatch -o -e . -i .go$$ -i Makefile -r . | xargs -n1 -I{}  make restart || make kill

restart: kill
	@make build
	@cd build; $(run_cmd) & echo $$! > $(pid)

kill:
	@kill `cat $(pid)` || true
