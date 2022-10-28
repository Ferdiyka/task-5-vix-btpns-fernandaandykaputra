package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/app"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/helpers"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/models"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authHelper helpers.AuthHelper
	jwtHelper  helpers.JWTHelper
}

func NewAuthController(authHelper helpers.AuthHelper, jwtHelper helpers.JWTHelper) AuthController {
	return &authController{
		authHelper: authHelper,
		jwtHelper:  jwtHelper,
	}
}

func (c *authController) Login(contx *gin.Context) {
	var registerData app.LoginApp
	errData := contx.ShouldBind(&registerData)
	if errData != nil {
		response := helpers.BuildErrorResponse("Maaf, gagal dalam memproses request", errData.Error(), helpers.EmptyObj{})
		contx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authFinal := c.authHelper.VerifyCredential(registerData.Email, registerData.Password)
	if v, ok := authFinal.(models.User); ok {
		generatedToken := c.jwtHelper.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helpers.BuildResponse(true, "OK", v)
		contx.JSON(http.StatusOK, response)
		return
	}

}

func (c *authController) Register(contx *gin.Context) {
	var registerData app.RegisterApp
	errData := contx.ShouldBind(&registerData)
	if errData != nil {
		response := helpers.BuildErrorResponse("Maaf, gagal dalam memproses request", errData.Error(), helpers.EmptyObj{})
		contx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authHelper.IsDuplicateEmail(registerData.Email) {
		response := helpers.BuildErrorResponse("Maaf, gagal dalam memproses request", "Duplicate email", helpers.EmptyObj{})
		contx.JSON(http.StatusConflict, response)
	} else {
		createUser := c.authHelper.CreateUser(registerData)
		token := c.jwtHelper.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		response := helpers.BuildResponse(true, "OK!", createUser)
		contx.JSON(http.StatusCreated, response)
	}
}