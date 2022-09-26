protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto
screen -dmS server go run server.go
#Open another terminal
screen -S client go run client.go
screen -r server
