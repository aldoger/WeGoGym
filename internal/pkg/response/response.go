package response

import (
	myerror "go-kpl/internal/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"-"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	Error      any    `json:"error,omitempty"`
	Data       any    `json:"data,omitempty"`
}

func NewSuccess(msg string, data any, meta ...any) Response {
	res := Response{
		StatusCode: 200,
		Success:    true,
		Message:    msg,
		Data:       data,
	}

	return res
}

func NewFailed(msg string, err error, meta ...any) Response {
	res := Response{
		StatusCode: http.StatusInternalServerError,
		Success:    false,
		Message:    msg,
		Error:      err.Error(),
	}

	if myErr, ok := err.(myerror.Error); ok {
		res.StatusCode = myErr.StatusCode
	}

	return res
}

func (r Response) Send(ctx *gin.Context) {
	sendStatus := r.StatusCode
	ctx.JSON(sendStatus, r)
}

func (r Response) SendWithAbort(ctx *gin.Context) {
	sendStatus := r.StatusCode
	ctx.AbortWithStatusJSON(sendStatus, r)
}
