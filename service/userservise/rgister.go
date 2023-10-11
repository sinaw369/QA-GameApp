package userservise

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/entity"
	"fmt"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {

	//TODO : replace md5
	//create new usr in storage
	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    GetMd5(req.Password),
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return create new user
	return dto.RegisterResponse{User: dto.UserInfo{
		ID:          createdUser.ID,
		Name:        createdUser.Name,
		PhoneNumber: createdUser.PhoneNumber,
	}}, nil

}
