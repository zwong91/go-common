package service

import (
	"context"
	"fmt"

	pb "go-common/app/tool/protoc-gen-bm/examples/helloworld/api"
)

// Service .
type Service struct{}

var _ pb.HelloServer = &Service{}

// SayHello .
func (s *Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("hello %s", req.Name),
	}, nil
}

// Echo .
func (s *Service) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{Content: req.Content}, nil
}
