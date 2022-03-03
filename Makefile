create:
	protoc --proto_path=wallet wallet/*.proto --go_out=gen/ 
	protoc --proto_path=wallet wallet/*.proto --go-grpc_out=gen/

clean:
	rm gen/wallet/*.go