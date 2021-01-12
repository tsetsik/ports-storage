# Starts the actual service
start:
	@echo 'Starting ports-storage'
	@docker-compose up -d

# Stops all the services
stop:
	@echo 'Stopping ports-storage'
	@docker-compose stop

# It executes the tests
test: start
	@echo 'Performing tests'
	@go test ./...