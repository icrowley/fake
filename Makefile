.DEFAULT_GOAL  = all
NAME           = fake
PACKAGE        = github.com/icrowley/$(NAME)
NUMCPUS        = $(shell cat /proc/cpuinfo | grep '^processor\s*:' | wc -l)


.PHONY: all
all: tools

.PHONY: $(NAME)
$(NAME):
	govendor remove +u
	govendor add +e
	govendor sync

.PHONY: build
build: $(NAME)


.PHONY: test
test: tools
	go test -v ./...

.PHONY: lint
lint: tools
	go vet ./...
	gometalinter                     \
		--deadline=5m            \
		--concurrency=$(NUMCPUS) \
		--exclude="(^|/)vendor/" \
		./...

.PHONY: check
check: lint test

.PHON: tools
tools:
	if [ ! -e "$(GOPATH)"/src/"github.com/kardianos/govendor" ]; then go get github.com/kardianos/govendor; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/rogpeppe/godef" ]; then go get github.com/rogpeppe/godef; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/nsf/gocode" ]; then go get github.com/nsf/gocode; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/stretchr/testify/assert" ]; then go get github.com/stretchr/testify/assert; fi
	if [ ! -e "$(GOPATH)"/src/"github.com/alecthomas/gometalinter" ]; then go get github.com/alecthomas/gometalinter && gometalinter --install; fi
