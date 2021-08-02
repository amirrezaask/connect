package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func OK(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, data)
}
