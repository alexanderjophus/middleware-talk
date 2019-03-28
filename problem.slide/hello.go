package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/trelore/middleware-talk/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// START SERVER OMIT
type S struct {
}

func (s S) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	if len(in.GetName()) > 8 { // HL1
		err := fmt.Sprintf("the name is too long: %s", in.GetName()) // HL1
		return nil, status.Error(codes.Internal, err)                // HL1
	} // HL1
	response := &pb.HelloResponse{
		Greeting: fmt.Sprintf("Hello, %s!", in.Name),
	}
	return response, nil
}

// END SERVER OMIT

// START OMIT
func main() {
	serverAddress := "127.0.0.1:8900"
	grpcServer := grpc.NewServer()
	s := S{}
	pb.RegisterGreetingServer(grpcServer, s)
	l, _ := net.Listen("tcp", serverAddress)
	go grpcServer.Serve(l)

	conn, _ := grpc.Dial(serverAddress, grpc.WithInsecure())
	c := pb.NewGreetingClient(conn)

	resp, err := c.Hello(context.Background(), &pb.HelloRequest{Name: "Alex"})
	fmt.Printf("response: %v\n", resp)
	fmt.Printf("error:    %v\n", err)
}

// END OMIT

// l, err := net.Listen("tcp", serverAddress)
// if err != nil {
// 	fmt.Printf("can't listen to %s: %v", serverAddress, err)
// 	return
// }
// defer l.Close()
// defer grpcServer.GracefulStop()

// conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
// if err != nil {
// 	fmt.Printf("can't dial %s: %v", serverAddress, err)
// 	return
// }
