package auth

import (
	"net/http"

	"github.com/adhtanjung/trackerstache_api/pkg/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register() gin.HandlerFunc
	// Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type AuthHandlerImpl struct {
	repository Repository
}

func NewAuthHandler(repository Repository) AuthHandler {
	return &AuthHandlerImpl{
		repository,
	}
}

func (a *AuthHandlerImpl) Register() gin.HandlerFunc {
	// return a.repository.Register(ctx, req)

	return func(ctx *gin.Context) {
		var user models.User
		ctx.BindJSON(&user)
		err := a.repository.Register(user)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	}
}
