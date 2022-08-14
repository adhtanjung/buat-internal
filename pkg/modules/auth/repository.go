package auth

import (
	"errors"

	"github.com/adhtanjung/trackerstache_api/pkg/db"
	"github.com/adhtanjung/trackerstache_api/pkg/models"
	"github.com/adhtanjung/trackerstache_api/pkg/utils"
)

type Repository interface {
	Register(auth models.RegisterRequest) error
	Login(body models.LoginRequest) (models.User, error)
}

type repository struct {
	H *db.Handler
}

func NewRepository(h *db.Handler) Repository {
	return &repository{H: h}
}

func (r *repository) Register(auth models.RegisterRequest) error {
	var user models.User

	if result := r.H.DB.Where(models.User{Email: auth.Email}).First(&user); result.Error == nil {
		return errors.New("email already exists")
	}

	user.Email = auth.Email
	user.Password = utils.HashPassword(auth.Password)

	r.H.DB.Create(&user)

	return nil
}

func (r *repository) Login(body models.LoginRequest) (models.User, error) {
	var user models.User
	db := r.H.DB
	if result := db.Where("email = ?", body.Email).First(&user); result.Error != nil {
		return user, errors.New("user not found")
	}
	passCheck := utils.CheckPasswordHash(body.Password, user.Password)

	if !passCheck {
		return user, errors.New("incorrect password ")
	}

	return user, nil
}
