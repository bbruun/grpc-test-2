.DEFAULT_GOAL = run

go-mod-tidy:
	go mod tidy

build-grpc: go-mod-tidy

	@echo "Building protobuf"
	@protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/server.proto

build: go-mod-tidy
	go build .

run: build
	go run . grpc-test-2

test:
	echo "Testing /minion.MinionService/RegisterMinion"
	echo '{"Name":"minion1"}' | grpc-client-cli --insecure --address localhost:4505 --service SubscriberService --method Subscribe
	echo '{"Name":"minion2"}' | grpc-client-cli --insecure --address localhost:4505 --service SubscriberService --method Subscribe
	# echo '{"readytoreceive":false,"result":"","success":false}' | grpc-client-cli --insecure --address localhost:4505 --service minion.MinionService --method GetCommands

