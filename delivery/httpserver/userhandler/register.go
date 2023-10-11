package userhandler

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userRegister(ctx echo.Context) error {
	var Req dto.RegisterRequest
	if err := ctx.Bind(&Req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if fieldErr, err := h.uservalidator.ValidateRegisterRequest(Req); err != nil {
		msg, code := httpmsg.Error(err)
		return ctx.JSON(code, echo.Map{
			"message": msg,
			"error":   fieldErr,
		})
	}
	resp, err := h.userSvc.Register(Req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusCreated, resp)
}
