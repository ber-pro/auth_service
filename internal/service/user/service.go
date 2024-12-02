package user

import (
	"auth/internal/converter"
	"auth/internal/repository"
	desc "auth/pkg/user_v1"
	"context"
	"log"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get(ctx context.Context, id int64) (*desc.User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		log.Printf("Service: Error while getting user with id=%d: %v", id, err)
		return nil, err
	}

	return converter.ModelUserToUser(user), nil
}

func (s *Service) Create(ctx context.Context, user *desc.UserInfo) (int64, error) {
	id, err := s.repo.Create(ctx, converter.UserInfoToModeUserInfo(user))
	if err != nil {
		log.Printf("Service: Error while creating user: %v", err)
		return 0, err
	}
	return id, nil
}

func (s *Service) Update(ctx context.Context, id int64, user *desc.UserInfo) error {
	err := s.repo.Update(ctx, id, converter.UserInfoToModeUserInfo(user))
	if err != nil {
		log.Printf("Service: Error while updating user with id=%d: %v", id, err)
		return err
	}
	log.Printf("Service: User with id=%d, update successful", id)
	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Printf("Service: Error while deleting user with id=%d: %v", id, err)
	}
	log.Printf("Service: User with id=%d, delete successful", id)
	return nil
}
