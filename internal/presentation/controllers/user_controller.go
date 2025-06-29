package controllers

import (
	"go-kpl/internal/application/dto"
	"go-kpl/internal/application/services"
	myerror "go-kpl/internal/pkg/errors"
	"go-kpl/internal/pkg/response"
	"net/http"
	"os"
	"time"

	qrcode "github.com/skip2/go-qrcode"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		GetMe(ctx *gin.Context)
		GenerateQrMe(ctx *gin.Context)
		Logout(ctx *gin.Context)
	}

	userController struct {
		userService services.UserService
	}
)

const MAX_AGE = 259200 // 3 hari

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var req dto.UserRegistrationDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	user, err := c.userService.Register(ctx, req)
	if err != nil {
		response.NewFailed("failed to register", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("registration success", user).Send(ctx)
}

func (c *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginDto
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	user, err := c.userService.Login(ctx, req)
	if err != nil {
		response.NewFailed("failed to login", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	setCustomCookie(ctx, "id", user.Id, MAX_AGE)
	setCustomCookie(ctx, "email", user.Email, MAX_AGE)
	setCustomCookie(ctx, "role", user.Role, MAX_AGE)
	setCustomCookie(ctx, "username", user.Username, MAX_AGE)

	response.NewSuccess("login successfully", user).Send(ctx)
}

func (c *userController) GetMe(ctx *gin.Context) {

	userId := ctx.Param("id")
	if userId == "" {
		response.NewFailed("Id user not found", myerror.New("user id not provided", http.StatusBadRequest)).Send(ctx)
		return
	}
	user, err := c.userService.GetMeDataById(ctx, userId)
	if err != nil {
		response.NewFailed("failed to retrive user data", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
	}

	response.NewSuccess("data successfuly retrive", user).Send(ctx)
}

func (c *userController) GenerateQrMe(ctx *gin.Context) {
	UserId, err := ctx.Cookie("id")
	if err != nil {
		response.NewFailed("user id is not found", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	endpoint := os.Getenv("BASE_URL") + "/api/user-membership/search-membership/" + UserId

	pngQR, err := qrcode.Encode(endpoint, qrcode.Medium, 256)
	if err != nil {
		response.NewFailed("failed to generate QR code", myerror.New(err.Error(), http.StatusInternalServerError)).Send(ctx)
		return
	}

	ctx.Data(http.StatusOK, "image/png", pngQR)
}

func (c *userController) Logout(ctx *gin.Context) {
	ctx.SetCookie("id", "", -1, "/", "", false, true)
	ctx.SetCookie("email", "", -1, "/", "", false, true)
	ctx.SetCookie("role", "", -1, "/", "", false, true)
	ctx.SetCookie("username", "", -1, "/", "", false, true)

	response.NewSuccess("logout successfully", nil).Send(ctx)
}

func setCustomCookie(ctx *gin.Context, name, value string, maxAge int) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		Domain:   "",
		MaxAge:   maxAge,
		Secure:   true, // wajib true untuk SameSite=None
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
	}

	http.SetCookie(ctx.Writer, &cookie)
}
