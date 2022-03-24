package middleware

import (
	"ArtistHome/internal/global"
	"ArtistHome/internal/model"
	"ArtistHome/internal/request"
	"ArtistHome/internal/response"
	"ArtistHome/internal/service"
	"ArtistHome/internal/util"
	"errors"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Jwt() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: global.AuthId,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					global.AuthId:   v.ID,
					global.AuthName: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &response.UserResult{ID: uint(claims[global.AuthId].(float64)), Name: claims[global.AuthName].(string)}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals request.Login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			user := new(service.User).SingleByName(loginVals.Username)

			if user != nil {
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password))
				if err == nil {
					return user, nil
				}
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			v, ok := data.(*response.UserResult)
			user := new(service.User).Single(v.ID)
			if ok && user != nil {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			util.NewError(c, code, errors.New(message))
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			util.NewResponse(c, http.StatusOK, "OK", gin.H{
				"accessToken": token,
				"expires":     expire.Unix(),
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			util.NewResponse(c, http.StatusOK, "OK", nil)
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			util.NewResponse(c, http.StatusOK, "OK", gin.H{
				"accessToken": token,
				"expires":     expire.Unix(),
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		global.LOGGER.Fatal("JWT Error#" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		global.LOGGER.Fatal("authMiddleware.MiddlewareInit() Error#" + errInit.Error())
	}
	return authMiddleware
}
