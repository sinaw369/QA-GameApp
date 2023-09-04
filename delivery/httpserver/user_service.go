package httpserver

import (
	"Q/A-GameApp/service/userservise"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userRegister(ctx echo.Context) error {
	var Req userservise.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	resp, err := s.userSvc.Register(Req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, resp)
}
func (s Server) userLogin(ctx echo.Context) error {
	var req userservise.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	response, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, response)
}
func (s Server) Profile(ctx echo.Context) error {
	authToken := ctx.Request().Header.Get("Authorization")
	claims, err := s.authSvc.VarifyToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	resp, err := s.userSvc.Profile(userservise.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, resp)
}
