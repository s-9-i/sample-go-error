include .env

.PHONY: build
build:
	CGO_ENABLED=0 go build -o api -trimpath -ldflags '-s -w' main.go

.PHONY: run
run:
	SENTRY_PUBLIC_DSN=${SENTRY_PUBLIC_DSN} ./api
