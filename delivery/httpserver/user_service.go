package httpserver

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/httpmsg"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userRegister(ctx echo.Context) error {
	var Req dto.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if fieldErr, err := s.uservalidator.ValidateRegisterRequest(Req); err != nil {
		msg, code := httpmsg.Error(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErr,
		})
	}
	resp, err := s.userSvc.Register(Req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, resp)
}
func (s Server) userLogin(ctx echo.Context) error {
	var req dto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	response, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (s Server) userProfile(ctx echo.Context) error {
	authToken := ctx.Request().Header.Get("Authorization")
	fmt.Println(authToken)
	claims, err := s.authSvc.VerifyToken(authToken)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	resp, err := s.userSvc.Profile(dto.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusOK, resp)
}
