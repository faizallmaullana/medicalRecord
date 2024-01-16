package auth

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

type GetInput struct {
	Nip      string `json:"nip"`
	Password string `json:"password"`
}

// LOGIN
// POST /api/v1/resource/login
func Login(c *gin.Context) {
	var input GetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the admin by username
	var user models.Users
	if err := models.DB.Where("nip = ?", input.Nip).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "NIP not found"})
		return
	}

	if input.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
