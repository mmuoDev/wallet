OUTPUT = main 
SERVICE_NAME = transactions

create:
	protoc --proto_path=wallet wallet/*.proto --go_out=gen/ 
	protoc --proto_path=wallet wallet/*.proto --go-grpc_out=gen/

clean:
	rm gen/wallet/*.go

build-local:
	go build -o $(OUTPUT) main.go


test:
	go test ./...

run: build-local
	@echo ">> Running application ..."
	DB_PORT=3306 \
	DB_HOST=localhost \
	DB_USER=root \
	DB_PASSWORD=password \
	DB_NAME=core \
	APP_PORT=9000 \
	./$(OUTPUT)