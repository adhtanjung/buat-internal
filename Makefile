server:
	go run cmd/main.go
swagger:
	swag init -g cmd/main.go --output docs/trackerstache --parseDependency --parseInternal
run-dev:
	swag init -g cmd/main.go --output docs/trackerstache --parseDependency=true --parseInternal
	go run cmd/main.go