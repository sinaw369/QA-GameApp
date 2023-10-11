package userhandler

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userLogin(ctx echo.Context) error {
	var req dto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if fieldErr, err := h.uservalidator.ValidateLoginRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErr,
		})
	}
	response, err := h.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, response)
}
