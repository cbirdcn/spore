package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"spore/internal/biz"

	pb2 "spore/api/bye/v1"
	pb "spore/api/helloworld/v1"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
	log *log.Helper
}

func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc : uc, log: log.NewHelper(logger)}
}

func (s *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	// something
	return &pb.HelloReply{Message: "Hello"}, nil
}

func (s *GreeterService) SayBye(ctx context.Context, req *pb2.ByeRequest) (*pb2.ByeReply, error) {
	return &pb2.ByeReply{Message : "Bye"}, nil
}
