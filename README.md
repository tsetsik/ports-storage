# ports-storage
Sample server application that supports grpc calls for storing ports data

Supported features:
- It does have client in it
- It can be started via docker compose `docker-compose up`
- It does have makefile and three tasks: (`start`, `stop` and `test`)

Supported grpc messages:
- For more information, look at the `storage.proto` file