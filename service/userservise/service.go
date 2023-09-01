package userservise

import (
	"Q/A-GameApp/entity"
	"Q/A-GameApp/pkg/phoneNumber"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}
type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO: we should verify phone number by verification code
	//validate phoneNumber
	if !phoneNumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}
	//check uniqueness of phone number

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique ")
		}
	}

	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be greater than 3")
	}
	// validate password
	// TODO: check password with regex pattern
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password length should be greater than 8")
	}
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
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return create new user
	return RegisterResponse{
		User: createdUser,
	}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginResponse struct {
}

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	//check the existing of phone number from repository
	//get user by phone number
	// TODO: it would be better to user two separate methods for existence
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	if user.Password != GetMd5(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	//compare user.password whit the req.password

	//return ok
	return LoginResponse{}, nil

}
func GetMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
