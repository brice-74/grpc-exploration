package main

import (
	"context"
	"log"
	"time"

	pb "github.com/brice-74/grpc-exploration/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewExplorationServiceClient(conn)

	unaryCall(client)
	serverStreamingCall(client)
	clientStreamingCall(client)
	bidirectionalStreamingCall(client)
}

func unaryCall(client pb.ExplorationServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.UnaryCall(ctx, &pb.ExplorationRequest{Message: "Hello Unary"})
	if err != nil {
		log.Fatalf("could not call Unary: %v", err)
	}
	log.Printf("Unary Response: %s", res.Response)
}

func serverStreamingCall(client pb.ExplorationServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.ServerStreamingCall(ctx, &pb.ExplorationRequest{Message: "Hello Server Streaming"})
	if err != nil {
		log.Fatalf("could not call Server Streaming: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("Server Streaming Response: %s", res.Response)
	}
}

func clientStreamingCall(client pb.ExplorationServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.ClientStreamingCall(ctx)
	if err != nil {
		log.Fatalf("could not call Client Streaming: %v", err)
	}
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.ExplorationRequest{Message: "Hello Client Streaming"}); err != nil {
			log.Fatalf("could not send message: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}
	log.Printf("Client Streaming Response: %s", res.Response)
}

func bidirectionalStreamingCall(client pb.ExplorationServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := client.BidirectionalStreamingCall(ctx)
	if err != nil {
		log.Fatalf("could not call Bidirectional Streaming: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				break
			}
			log.Printf("Bidirectional Streaming Response: %s", res.Response)
		}
		close(waitc)
	}()
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.ExplorationRequest{Message: "Hello Bidirectional Streaming"}); err != nil {
			log.Fatalf("could not send message: %v", err)
		}
	}
	stream.CloseSend()
	<-waitc
}
