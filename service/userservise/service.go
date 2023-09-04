package userservise

import (
	"Q/A-GameApp/entity"
	"Q/A-GameApp/pkg/phoneNumber"
	"Q/A-GameApp/pkg/richerror"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}
type Service struct {
	repo Repository
	auth AuthGenerator
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type UserInfo struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
type RegisterResponse struct {
	User UserInfo `json:"user"`
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{auth: authGenerator, repo: repo}
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
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
	return RegisterResponse{UserInfo{
		ID:          createdUser.ID,
		Name:        createdUser.Name,
		PhoneNumber: createdUser.PhoneNumber,
	}}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type LoginResponse struct {
	User   UserInfo `json:"user"`
	Tokens Token    `json:"token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	const op = "userService.login"
	//check the existing of phone number from repository
	//get user by phone number
	// TODO: it would be better to user two separate methods for existence
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).WhitWarpError(err).WhitMeta(map[string]interface{}{"Phone_number": req.PhoneNumber})
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	if user.Password != GetMd5(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}

	//compare user.password whit the req.password
	//---//
	// jwt token
	AccessToken, err := s.auth.CreateAccessToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	RefreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return ok
	return LoginResponse{
		User: UserInfo{
			ID:          user.ID,
			PhoneNumber: user.PhoneNumber,
			Name:        user.Name,
		},
		Tokens: Token{
			RefreshToken: RefreshToken,
			AccessToken:  AccessToken,
		},
	}, nil

}
func GetMd5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}
type ProfileResponse struct {
	Name string `json:"name"`
}

// All Requests inputs for interactor / service should sanitize

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	const op = "userService.profile"
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, richerror.New(op).
			WhitWarpError(err).
			WhitMeta(map[string]interface{}{"request": req})
	}
	return ProfileResponse{Name: user.Name}, nil
}
