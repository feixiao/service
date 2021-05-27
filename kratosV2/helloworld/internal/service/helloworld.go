
package service

import(
	"context"

	pb "helloworld/api/helloworld"
)

type HelloworldService struct {
	pb.UnimplementedHelloworldServer
}

func NewHelloworldService() *HelloworldService {
	return &HelloworldService{}
}

func (s *HelloworldService) CreateHelloworld(ctx context.Context, req *pb.CreateHelloworldRequest) (*pb.CreateHelloworldReply, error) {
	return &pb.CreateHelloworldReply{}, nil
}
func (s *HelloworldService) UpdateHelloworld(ctx context.Context, req *pb.UpdateHelloworldRequest) (*pb.UpdateHelloworldReply, error) {
	return &pb.UpdateHelloworldReply{}, nil
}
func (s *HelloworldService) DeleteHelloworld(ctx context.Context, req *pb.DeleteHelloworldRequest) (*pb.DeleteHelloworldReply, error) {
	return &pb.DeleteHelloworldReply{}, nil
}
func (s *HelloworldService) GetHelloworld(ctx context.Context, req *pb.GetHelloworldRequest) (*pb.GetHelloworldReply, error) {
	return &pb.GetHelloworldReply{}, nil
}
func (s *HelloworldService) ListHelloworld(ctx context.Context, req *pb.ListHelloworldRequest) (*pb.ListHelloworldReply, error) {
	return &pb.ListHelloworldReply{}, nil
}
