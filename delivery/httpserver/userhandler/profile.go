package userhandler

import (
	"Q/A-GameApp/dto"
	"Q/A-GameApp/pkg/httpmsg"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userProfile(ctx echo.Context) error {
	authToken := ctx.Request().Header.Get("Authorization")
	fmt.Println(authToken)
	claims, err := h.authSvc.VerifyToken(authToken)
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	resp, err := h.userSvc.Profile(dto.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}
	return ctx.JSON(http.StatusOK, resp)
}
