package userhandler

import (
	"Q/A-GameApp/service/authservice"
	"Q/A-GameApp/service/userservise"
	"Q/A-GameApp/validator/uservalidator"
)

type Handler struct {
	authSvc       authservice.Service
	userSvc       userservise.Service
	uservalidator uservalidator.Validator
}

func New(authSvc authservice.Service, userSvc userservise.Service, uservalidator uservalidator.Validator) Handler {
	return Handler{
		authSvc:       authSvc,
		userSvc:       userSvc,
		uservalidator: uservalidator,
	}
}
