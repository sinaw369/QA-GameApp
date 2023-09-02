package httpserver

import (
	"Q/A-GameApp/service/userservise"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (s Server) userRegister(ctx echo.Context) error {
	var uReq userservise.RegisterRequest
	if err := ctx.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	resp, err := s.userSvc.Register(uReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, resp)
}
