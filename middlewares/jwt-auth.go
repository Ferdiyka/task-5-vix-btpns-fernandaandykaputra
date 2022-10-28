package middlewares

import (
	"log"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/Ferdiyka/task-5-vix-btpns-fernandaandykaputra/helpers"
)

func AuthJWT(jwtHelper helpers.JWTHelper) gin.HandlerFunc {
	return func(a *gin.Context) {
		authHeader := a.GetHeader("Authorization")
		if authHeader == "" {
			response := helpers.BuildErrorResponse("Maaf, gagal memproses request", "Token tidak ditemukan", nil)
			a.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtHelper.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helpers.BuildErrorResponse("Token tidak valid", err.Error(), nil)
			a.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}