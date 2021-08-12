package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
    Success bool
    Message string `json:"message,omitempty"`
    Payload interface{} `json:"payload,omitempty"`
}


func Response(ctx echo.Context, status int, r *APIResponse) error {
	return ctx.JSON(status, r)
}

func OK(ctx echo.Context, data ...interface{}) error {
    var toSend interface{} 
    if len(data) == 1 {
        toSend = data[0]
    } else {
        toSend = data
    }
    return Response(ctx, http.StatusOK, &APIResponse{
        Success: true,
        Payload: toSend,
    })
}

func ServerErr(ctx echo.Context, err error) error {
    return Response(ctx, http.StatusInternalServerError, &APIResponse {
        Success: false,
        Message: err.Error(),
    })
}

func ClientErr(ctx echo.Context, err error) error {
    return Response(ctx, http.StatusBadRequest, &APIResponse {
        Success: false,
        Message: err.Error(),
    })
}
