package auth

import (
	"net/http"

	"github.com/adhtanjung/trackerstache_api/pkg/utils"
	"github.com/gin-gonic/gin"
	//add binding import
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

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Accept json
// @Produce json
// @Param body body RegisterRequest true "Register Request"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} RegisterResponse
// @Failure 500 {object} RegisterResponse
// @Router /auth/register [post]
func (a *AuthHandlerImpl) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body RegisterRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			// var errorMessage string
			// if s := utils.CustomValidateStruct(err); s != "" {
			// 	errorMessage = utils.CustomValidateStruct(err)
			// } else {
			// 	errorMessage = err.Error()
			// }
			ctx.JSON(http.StatusBadRequest, RegisterResponse{
				Status: http.StatusBadRequest,
				Error:  utils.CustomValidateStruct(err),
			})
			return
		}
		err := a.repository.Register(body)

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
