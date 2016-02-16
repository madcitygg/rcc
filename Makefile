.PHONY: release

version = $(shell sed -n 's/.*version = "\([^"]*\)\s*"/\1/p' main.go)

help:
	@echo "release - build binaries for current version (${version})"

release:
	mkdir -p release/
	rm release/*

	GOOS=windows GOARCH=amd64 go build -o release/rcc.exe . && \
		cd release && \
		zip -X -o rcc_${version}_windows_386.zip rcc.exe && \
		rm rcc.exe

	GOOS=darwin GOARCH=amd64 go build -o release/rcc . && \
		cd release && \
		zip -X -o rcc_${version}_darwin_amd64.zip rcc && \
		rm rcc

	GOOS=linux GOARCH=amd64 go build -o release/rcc . && \
		cd release && \
		zip -X -o rcc_${version}_linux_amd64.zip rcc && \
		rm rcc
