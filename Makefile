GIT_COMMIT:=$(shell git describe --dirty --always)
VERSION:=$(shell grep 'VERSION' pkg/version/version.go | awk '{ print $$4 }' | tr -d '"')

run:
	go run -ldflags "-s -w -X github.com/heover1cks/pgen/pkg/version.REVISION=$(GIT_COMMIT)" main.go password -L 24

build:
	GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/heover1cks/pgen/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/pgen ./main.go
	cp ./bin/pgen /usr/local/bin
build-linux:
	GOARCH=amd64 GOOS=linux GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/heover1cks/pgen/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/pgen-linux-amd64-$(VERSION) ./main.go
build-darwin:
	GOARCH=amd64 GOOS=darwin GIT_COMMIT=$$(git rev-list -1 HEAD) && CGO_ENABLED=0 go build  -ldflags "-s -w -X github.com/heover1cks/pgen/pkg/version.REVISION=$(GIT_COMMIT)" -a -o ./bin/pgen-darwin-amd64-$(VERSION) ./main.go
tidy:
	rm -f go.sum; go mod tidy -compat=1.19

vet:
	go vet ./...

fmt:
	gofmt -l -s -w ./
	goimports -l -w ./

release:
	git tag $(VERSION)
	git push origin $(VERSION)