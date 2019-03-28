package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/trelore/middleware-talk/proto"
	"google.golang.org/grpc"
)

// START SERVER OMIT
type S struct {
}

func (s S) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
		Greeting: fmt.Sprintf("Hello, %s!", in.Name),
	}
	return response, nil
}

// END SERVER OMIT

// START OMIT
func main() {
	serverAddress := "127.0.0.1:8900"        // HL1
	grpcServer := grpc.NewServer()           // HL1
	s := S{}                                 // HL1
	pb.RegisterGreetingServer(grpcServer, s) // HL1
	l, _ := net.Listen("tcp", serverAddress) // HL1
	go grpcServer.Serve(l)                   // HL1

	conn, _ := grpc.Dial(serverAddress, grpc.WithInsecure()) // HL2
	c := pb.NewGreetingClient(conn)                          // HL2

	resp, err := c.Hello(context.Background(), &pb.HelloRequest{Name: "Alexander"}) // HL3
	fmt.Printf("response: %v\n", resp)                                              // HL3
	fmt.Printf("error:    %v\n", err)                                               // HL3
}

// END OMIT

// interceptor := grpc.UnaryInterceptor(grpcmask.UnaryServerInterceptor())
// grpcServer := grpc.NewServer(interceptor)

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
