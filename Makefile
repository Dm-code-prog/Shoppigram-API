ifeq ($(MAKECMDGOALS),run)
include .env
export
endif
run	:
	go run cmd/api/main.go cmd/api/environment.go cmd/api/adapters.go

test:
	go test ./... -race

sqlc:
	for dir in ./internal/*/; do \
		(cd "$$dir" && if test -f sqlc.yaml; then \
			echo "running sqlc generate for $${dir} ...";\
			sqlc generate; \
		fi)\
	done

