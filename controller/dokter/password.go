package dokter

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

type IniInput struct {
	Password string `json:"password"`
}

// fungsi untuk merubah password
// POST /dokter/password/change/:id
func ChangePassword(c *gin.Context) {
	var password models.Users
	if err := models.DB.Where("id = ? ", c.Param("id")).First(&password).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access this"})
		return
	}

	// validate input
	var input IniInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := models.Users{
		Password: input.Password,
	}

	models.DB.Model(&data).Updates(data)
}
