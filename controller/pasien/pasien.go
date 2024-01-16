package pasien

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

func GetAllPasien(c *gin.Context) {
	var pasien []models.Patient
	if err := models.DB.Find(&pasien).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pasien})
}
