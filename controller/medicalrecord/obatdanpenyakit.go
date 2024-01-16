package medicalrecord

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

// GET /admin/get/all/obat/penyakit
func GetAllObatPenyakit(c *gin.Context) {
	var medicines []models.Medicine
	var diagnosis []models.Diagnosis

	if err := models.DB.Find(&medicines).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := models.DB.Find(&diagnosis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"medicines": medicines,
		"diagnosis": diagnosis,
	})
}
