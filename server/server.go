package server

import (
	"fmt"
	"time"

	proto "github.com/bbruun/grpc-test-2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type mySubscriberService struct {
	proto.UnimplementedSubscriberServiceServer
}

func registerSubscriberServics(grpcServer *grpc.Server) {
	reflection.Register(grpcServer) // enable reflection from client side
	service := &mySubscriberService{}
	proto.RegisterSubscriberServiceServer(grpcServer, service)
}

// func (m *mySubscriberService) Create(*proto.FromClient) (*proto.ToClient, error) {
// 	return &proto.ToClient{}, nil
// }

// func (s *mySubscriberService) Subscribe(*proto.FromClient, proto.SubscriberService_SubscribeServer) error {

// 	//TODO: make chan and send to global queue for later retrieval
// 	//TODO: setup select statement to return statements to the Subscirbed client
// 	return nil
// }

// func (m *mySubscriberService) Subscribe(ctx context.Context, input *proto.FromClient) (*proto.ToClient, error) {
func (m *mySubscriberService) Subscribe(fromClient *proto.FromClient, toClient proto.SubscriberService_SubscribeServer) error {

	a := time.Now().UTC()

	done := make(chan bool)
	_ = done

	fmt.Printf("On %d Subscribe() received: %+v\n", a.UnixNano(), fromClient)

	// return toclient
	for {
		b := time.Now().UTC()

		messageToReturn := fmt.Sprintf("This is the string the client will get %d", b.UnixNano())
		msg := proto.ToClient{
			Message: messageToReturn,
		}
		if err := toClient.SendMsg(&msg); err != nil {
			fmt.Printf("failed to send command to minion, closing server port: %s\n", err)
			toClient.Context().Done()
			return fmt.Errorf("client is not connected")
		}
		time.Sleep(2 * time.Second)

	}

	return nil
}
