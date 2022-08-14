package auth

import (
	"net/http"

	"github.com/adhtanjung/trackerstache_api/pkg/models"
	"github.com/adhtanjung/trackerstache_api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Register() gin.HandlerFunc
	Login() gin.HandlerFunc
	// Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error)
}

type AuthHandlerImpl struct {
	repository Repository
	Jwt        utils.JwtWrapper
}

func NewAuthHandler(repository Repository) AuthHandler {
	return &AuthHandlerImpl{
		repository,
		utils.JwtWrapper{},
	}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Accept json
// @Produce json
// @Param body body models.RegisterRequest true "Register Request"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.RegisterResponse
// @Failure 500 {object} models.RegisterResponse
// @Router /auth/register [post]
func (a *AuthHandlerImpl) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.RegisterRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, models.RegisterResponse{
				Status: http.StatusBadRequest,
				Error:  utils.CustomValidateStruct(err),
			})
			return
		}
		err := a.repository.Register(body)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				models.RegisterResponse{
					Status: http.StatusInternalServerError,
					Error:  utils.CustomValidateStruct(err),
				})
			return
		}

		ctx.JSON(http.StatusOK, models.RegisterResponse{
			Status:  http.StatusOK,
			Message: "User registered successfully",
		})

	}
}

// Login godoc
// @Summary User login
// @Description login
// @Accept json
// @Produce json
// @Param body body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.WebResponse
// @Failure 500 {object} models.WebResponse
// @Router /auth/login [post]
func (a *AuthHandlerImpl) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body models.LoginRequest
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, models.RegisterResponse{
				Status: http.StatusBadRequest,
				Error:  utils.CustomValidateStruct(err),
			})
			return
		}
		user, err := a.repository.Login(body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				models.RegisterResponse{
					Status: http.StatusInternalServerError,
					Error:  utils.CustomValidateStruct(err),
				})
			return
		}
		token, _ := a.Jwt.GenerateToken(user)
		ctx.JSON(http.StatusOK, models.LoginResponse{
			Status: http.StatusOK,
			Token:  token,
		})

	}
}
