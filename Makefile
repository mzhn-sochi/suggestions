proto:
	protoc -I ./proto ./proto/*.proto  --go_out=. --go-grpc_out=.
