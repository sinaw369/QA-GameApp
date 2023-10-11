package userservise

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/richerror"
	"fmt"
)

func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	const op = "userService.login"
	//check the existing of phone number from repository
	//get user by phone number
	// TODO: it would be better to user two separate methods for existence
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WhitWarpError(err).WhitMeta(map[string]interface{}{"Phone_number": req.PhoneNumber})
	}
	if !exist {
		return dto.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	if user.Password != GetMd5(req.Password) {
		return dto.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	//compare user.password whit the req.password
	//---//
	// jwt token
	AccessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	RefreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return ok
	return dto.LoginResponse{
		User: dto.UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: dto.Token{
			RefreshToken: RefreshToken,
			AccessToken:  AccessToken,
		},
	}, nil

}
