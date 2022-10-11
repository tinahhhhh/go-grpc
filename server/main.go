package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/tinahhhhh/go-grpc/spam"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement spam.AssessmentServer.
type server struct {
	pb.UnimplementedAssessmentServer
}

// CheckSpam implements spam.AssessmentServer.
func (s *server) CheckSpam(ctx context.Context, in *pb.AssessmentRequest) (*pb.AssessmentReply, error) {
	log.Printf("Received: %v", in.GetEntity())
	return &pb.AssessmentReply{Message: "Result: " + in.GetEntity()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAssessmentServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
