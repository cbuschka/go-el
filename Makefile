PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ifeq (${GOPATH},)
        GOPATH := ${HOME}/go
endif

test:
	@cd $(PROJECT_DIR) && \
	go test ./...

tidy:
	@cd $(PROJECT_DIR) && \
	go mod download && \
	go mod tidy

generate:
	@cd ${PROJECT_DIR}
	go get github.com/goccmack/gocc
	go install github.com/goccmack/gocc
	rm -rf ${PROJECT_DIR}/internal/generated/ && \
	mkdir -p ${PROJECT_DIR}/internal/generated/ && \
	${GOPATH}/bin/gocc -o internal/generated/ internal/grammar.bnf
