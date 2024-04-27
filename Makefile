ifeq ($(MAKECMDGOALS),run)
include .env
export
endif
run	:
	go run cmd/api/main.go cmd/api/environment.go