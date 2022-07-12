build:
	cd func && go build .
	cd pkg/orm && go build .
PHONY: build
