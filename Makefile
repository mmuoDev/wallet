OUTPUT = main 
SERVICE_NAME = wallet

build-local:
	go build -o $(OUTPUT) ./cmd/$(SERVICE_NAME)/main.go

clean:
	rm -f $(OUTPUT)

test:
	go test ./...

run: build-local
	@echo ">> Running application ..."
	DB_PORT= \
	DB_HOST= \
	DB_USER= \
	DB_PASSWORD= \
	DB_NAME=core \
	APP_PORT= \
	./$(OUTPUT)