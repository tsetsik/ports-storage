# Starts the actual service
start:
	@echo 'Starting ports-storage'
	@docker-compose up -d

# Stops all the services
stop:
	@echo 'Stopping ports-storage'
	@docker-compose stop

# Generates the go files using storage.proto file as source
proto:
	@echo 'Generating proto files'
	@protoc --go_out=plugins=grpc:internal/storage ./storage.proto

# It executes the tests
test: start
	@echo 'Performing tests'
	@go test ./...