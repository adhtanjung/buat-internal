server:
	go run cmd/main.go
swagger:
	swag init -g cmd/main.go --output docs/trackerstache --parseDependency --parseInternal
run-dev:
	swag init -g cmd/main.go --output docs/trackerstache --parseDependency --parseInternal
	go run cmd/main.go