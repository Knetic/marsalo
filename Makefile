default: build
all: package

export GOPATH=$(CURDIR)/
export GOBIN=$(CURDIR)/.temp/

init: clean
	go get ./...

build: init
	go build -o ./.output/httpMarshal .

test:
	go test
	go test -bench=.

clean:
	@rm -rf ./.output/

fmt:
	@go fmt .
	@go fmt ./src/httpMarshal

dist: build test

	export GOOS=linux; \
	export GOARCH=amd64; \
	go build -o ./.output/httpMarshal64 .

	export GOOS=linux; \
	export GOARCH=386; \
	go build -o ./.output/httpMarshal32 .

	export GOOS=darwin; \
	export GOARCH=amd64; \
	go build -o ./.output/httpMarshal_osx .

	export GOOS=windows; \
	export GOARCH=amd64; \
	go build -o ./.output/httpMarshal.exe .



package: dist

ifeq ($(shell which fpm), )
	@echo "FPM is not installed, no packages will be made."
	@echo "https://github.com/jordansissel/fpm"
	@exit 1
endif

ifeq ($(HTTPM_VERSION), )

	@echo "No 'HTTPM_VERSION' was specified."
	@echo "Export a 'HTTPM_VERSION' environment variable to perform a package"
	@exit 1
endif

	fpm \
		--log error \
		-s dir \
		-t deb \
		-v $(HTTPM_VERSION) \
		-n httpMarshal \
		./.output/httpMarshal64=/usr/local/bin/httpMarshal \
		./docs/httpMarshal.7=/usr/share/man/man7/httpMarshal.7 \
		./autocomplete/httpMarshal=/etc/bash_completion.d/httpMarshal

	fpm \
		--log error \
		-s dir \
		-t deb \
		-v $(HTTPM_VERSION) \
		-n httpMarshal \
		-a i686 \
		./.output/httpMarshal32=/usr/local/bin/httpMarshal \
		./docs/httpMarshal.7=/usr/share/man/man7/httpMarshal.7 \
		./autocomplete/httpMarshal=/etc/bash_completion.d/httpMarshal

	@mv ./*.deb ./.output/

	fpm \
		--log error \
		-s dir \
		-t rpm \
		-v $(HTTPM_VERSION) \
		-n httpMarshal \
		./.output/httpMarshal64=/usr/local/bin/httpMarshal \
		./docs/httpMarshal.7=/usr/share/man/man7/httpMarshal.7 \
		./autocomplete/httpMarshal=/etc/bash_completion.d/httpMarshal
	fpm \
		--log error \
		-s dir \
		-t rpm \
		-v $(HTTPM_VERSION) \
		-n httpMarshal \
		-a i686 \
		./.output/httpMarshal32=/usr/local/bin/httpMarshal \
		./docs/httpMarshal.7=/usr/share/man/man7/httpMarshal.7 \
		./autocomplete/httpMarshal=/etc/bash_completion.d/httpMarshal

	@mv ./*.rpm ./.output/
