package dokter

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

// ambil dokter berdasar id yang sedang login
// id dikirim dari front end
func GetDokter(c *gin.Context) {
	var dokter models.Users
	if err := models.DB.Preload("Profile").Where("id = ? ", c.Param("id")).First(&dokter).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access this"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dokter})
}

func GetDokterAll(c *gin.Context) {
	var dokter []models.Users
	if err := models.DB.Preload("Profile").Where("role = ? ", "dokter").Find(&dokter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dokter})
}
