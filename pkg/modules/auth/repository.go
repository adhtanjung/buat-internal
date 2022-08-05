package auth

import (
	"errors"
	"net/http"

	"github.com/adhtanjung/trackerstache_api/pkg/db"
	"github.com/adhtanjung/trackerstache_api/pkg/models"
	"github.com/adhtanjung/trackerstache_api/pkg/utils"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Status int32  `json:"status"`
	Error  string `json:"error"`
}
type Repository interface {
	Register(auth models.User) error
}

type repository struct {
	H *db.Handler
}

func NewRepository(h *db.Handler) Repository {
	return &repository{H: h}
}

func (r *repository) Register(auth models.User) error {
	var user models.User
	// log.Println(req)

	if result := r.H.DB.Where(models.User{Email: auth.Email}).First(&user); result.Error == nil {
		return utils.NewError(http.StatusBadRequest, errors.New("Email already exists"))
	}

	user.Email = auth.Email
	user.Password = utils.HashPassword(auth.Password)

	r.H.DB.Create(&user)

	return nil
}
