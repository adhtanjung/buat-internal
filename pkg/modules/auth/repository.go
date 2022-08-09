package auth

import (
	"errors"

	"github.com/adhtanjung/trackerstache_api/pkg/db"
	"github.com/adhtanjung/trackerstache_api/pkg/models"
	"github.com/adhtanjung/trackerstache_api/pkg/utils"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	Status int32               `json:"status"`
	Error  []map[string]string `json:"errors"`
}
type Repository interface {
	Register(auth RegisterRequest) error
}

type repository struct {
	H *db.Handler
}

func NewRepository(h *db.Handler) Repository {
	return &repository{H: h}
}

func (r *repository) Register(auth RegisterRequest) error {
	var user models.User

	if result := r.H.DB.Where(models.User{Email: auth.Email}).First(&user); result.Error == nil {
		return errors.New("email already exists")
	}

	user.Email = auth.Email
	user.Password = utils.HashPassword(auth.Password)

	r.H.DB.Create(&user)

	return nil
}
