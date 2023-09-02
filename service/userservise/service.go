package userservise

import (
	"Q/A-GameApp/entity"
	"Q/A-GameApp/pkg/phoneNumber"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}
type Service struct {
	repo    Repository
	signKey string
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	User entity.User
}

func New(repo Repository, signKey string) *Service {
	return &Service{repo: repo, signKey: signKey}
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
	AccessToken string `json:"access_token"`
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
	//---//
	// jwt token
	token, err := createToken(user.ID, s.signKey)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	//return ok
	return LoginResponse{AccessToken: token}, nil

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

func (s *Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	// getUserByID
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		// TODO : we can use rich error
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}
	return ProfileResponse{Name: user.Name}, nil
}

type Claims struct {
	RegisteredClaims jwt.RegisteredClaims
	UserID           uint
}

func (c Claims) Valid() error {
	return nil
}
func createToken(userID uint, signKey string) (string, error) {
	//create a new token
	// TODO : replace with rsa256 - https://github.com/golang-jwt/jwt/
	//see our claims
	Claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := accessToken.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
