# path:
# 	export PATH="$PATH:$(go env GOPATH)/bin"

gen:
	protoc --proto_path=proto proto/*.proto --go_out=pb --go_opt=paths=source_relative

gen_grpc:
	protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    --proto_path=proto proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go

client:
	go run cmd/client/main.go -address 0.0.0.0:3000

server:
	go run cmd/server/main.go -port 3000