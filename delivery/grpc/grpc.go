package grpc

import (
	"context"

	"github.com/JIeeiroSst/go-app/domain"
	"github.com/JIeeiroSst/go-app/proto"
)

type Service struct {
	repo domain.Repository
	proto.UnimplementedUserProfileServer
}

func NewService(repo domain.Repository) *Service {
	return &Service{
		repo:repo,
	}
} 

func (s *Service) UpdateEmail(ctx context.Context,req *proto.UpdateEmailRequest) (*proto.Response,error){
	profile:=domain.Profile{
		Email: req.GetEmail(),
	}
	id:=int(req.GetId())
	err := s.repo.UpdateProfile(id,profile)
	if err!=nil {
		return &proto.Response{
			Ok:      false,
			Message: "cannot update failed",
		}, err
	}
	return &proto.Response{
		Ok:      false,
		Message: "can update success ",
	}, nil
}

func (s *Service )CreateEmail(ctx context.Context,req *proto.CreateEmailRequest) (*proto.Response,error){
	profile:=domain.Profile {
		Name: req.GetName(),
		Email: req.GetEmail(),
		UserId: int(req.GetUserId()),
	}
	err:=s.repo.CreateProfile(profile)
	if err!=nil{
		if err!=nil {
			return &proto.Response{
				Ok:      false,
				Message: "cannot create failed",
			}, err
		}
	}
	return &proto.Response{
		Ok:      false,
		Message: "can create success ",
	}, nil
}

func (s *Service)DeleteEmail(ctx context.Context,req *proto.DeleteEmailRequest) (*proto.Response,error){
	id:=int(req.GetId())
	err:=s.repo.DeleteProfile(id)
	if err!=nil {
		return &proto.Response{
			Ok:      false,
			Message: "cannot delete failed",
		}, err
	}
	return &proto.Response{
		Ok:      false,
		Message: "can delete success ",
	}, nil
}