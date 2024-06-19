package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brice-74/grpc-exploration/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExplorationServiceServer
}

func (s *server) UnaryCall(ctx context.Context, req *pb.ExplorationRequest) (*pb.ExplorationResponse, error) {
	return &pb.ExplorationResponse{Response: "Unary response to: " + req.Message}, nil
}

func (s *server) ServerStreamingCall(req *pb.ExplorationRequest, stream pb.ExplorationService_ServerStreamingCallServer) error {
	for i := 0; i < 5; i++ {
		if err := stream.Send(&pb.ExplorationResponse{Response: "Streaming response " + req.Message}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) ClientStreamingCall(stream pb.ExplorationService_ClientStreamingCallServer) error {
	var message string
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		message += req.Message
	}
	return stream.SendAndClose(&pb.ExplorationResponse{Response: "Client streaming response to: " + message})
}

func (s *server) BidirectionalStreamingCall(stream pb.ExplorationService_BidirectionalStreamingCallServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		if err := stream.Send(&pb.ExplorationResponse{Response: "Bidirectional response to: " + req.Message}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExplorationServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
