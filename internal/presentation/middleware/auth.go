package middleware

import (
	myerror "go-kpl/internal/pkg/errors"
	"go-kpl/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MESSAGE_NOT_AUTHORIZE    = "user not authorized"
	MESSAGE_NOT_AUTHENTICATE = "user is not authenticate"
	MESSAGE_COOKIE_EMPTY     = "cookie is empty"
)

func (m Middleware) IsCookieEmpty(userId string, userEmail string, userRole string, userName string) bool {
	return userId == "" || userEmail == "" || userRole == "" || userName == ""
}

func (m Middleware) Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, err := ctx.Cookie("role")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing role cookie"})
			return
		}

		if role == "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "role checker not initialized"})
			return
		}

		ctx.Next()
	}
}

func (m Middleware) OnlyAllow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole, err := ctx.Cookie("role")
		if err != nil {
			response.NewFailed(MESSAGE_NOT_AUTHENTICATE, myerror.New(err.Error(), http.StatusUnauthorized)).SendWithAbort(ctx)
			return
		}

		isAdmin := m.UserRole.IsAdmin(userRole)
		if !isAdmin {
			response.NewFailed(MESSAGE_NOT_AUTHORIZE, myerror.New(MESSAGE_NOT_AUTHORIZE, http.StatusForbidden)).SendWithAbort(ctx)
			return
		}

		ctx.Next()
	}
}
