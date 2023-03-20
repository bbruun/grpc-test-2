package server

import (
	"fmt"
	"time"

	"github.com/bbruun/grpc-test-2/messaging"
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

	done := make(chan bool)
	_ = done
	var mc *messaging.Minions = messaging.MinionStateCollector

	msgch := make(chan any)
	mi := messaging.MinionInfo{
		Name:                  fromClient.Name,
		MessageFromClient:     fromClient.MessageFromClient,
		MessageToClient:       fromClient.MessageToClient,
		CommunicationsChannel: msgch,
		IsConnected:           true,
	}
	mc.AddMinion(&mi)

	// return toclient
	for {
		b := time.Now().UTC()

		messageToReturn := fmt.Sprintf("This is the string the client will get %d", b.UnixNano())
		msg := proto.ToClient{
			Message: messageToReturn,
		}
		if err := toClient.SendMsg(&msg); err != nil {
			mi.IsConnected = false
			fmt.Printf("minion %s disconected (%s)\n", mi.Name, err)
			toClient.Context().Done()
			return fmt.Errorf("client %s is not connected", mi.Name)
		}
		time.Sleep(2 * time.Second)

	}

	return nil
}
