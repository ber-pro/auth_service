package user

import (
	"auth/internal/service"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type API struct {
	desc.UnimplementedUserV1Server
	serviceLayer service.Service
}

func NewAPI(serviceLayer service.Service) *API {
	return &API{serviceLayer: serviceLayer}
}

func (a *API) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userInfo, err := a.serviceLayer.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	log.Printf("User with id=%d, get records[name: %s, email: %s]", req.GetId())

	return &desc.GetResponse{User: userInfo}, nil
}

func (a *API) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	var empty *emptypb.Empty
	err := a.serviceLayer.Update(ctx, req.GetId(), &desc.UserInfo{
		Name:  req.Name.GetValue(),
		Email: req.Email.GetValue(),
		Role:  req.Role,
	})

	if err != nil {
		log.Printf("Can't update user with id=%d", req.Id)
		return nil, err
	}

	log.Printf("User with id=%d, update records[name: %s, email: %s]", req.Id, req.Name, req.Email)
	return empty, nil
}

func (a *API) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userID, err := a.serviceLayer.Create(ctx, &desc.UserInfo{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Role:     req.GetRole(),
	})
	if err != nil {
		return nil, err
	}

	log.Printf("User with id=%d, created", userID)

	return &desc.CreateResponse{Id: userID}, nil
}

func (a *API) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {

	var empty *emptypb.Empty
	err := a.serviceLayer.Delete(ctx, req.Id)
	if err != nil {
		log.Printf("User with id=%d, delete", req.Id)
		return nil, err
	}

	log.Printf("User with id=%d, delete", req.Id)
	return empty, nil
}
