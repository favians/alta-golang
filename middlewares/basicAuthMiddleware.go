package middlewares

import (
	"go-training-restful/config"
	"go-training-restful/models"

	"github.com/labstack/echo"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}

	return true, nil
}
