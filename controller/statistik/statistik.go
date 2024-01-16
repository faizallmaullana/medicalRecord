package statistik

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

// get sum of medicine, patient, disease, and medical staff
// GET /statistics
func Statistics(c *gin.Context) {
	var totalMedicine int64
	var totalPatient int64
	var totalDiagnosis int64
	var totalMedicalStaff int64

	if err := models.DB.Model(&models.Medicine{}).Count(&totalMedicine).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := models.DB.Model(&models.Patient{}).Count(&totalPatient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := models.DB.Model(&models.Diagnosis{}).Count(&totalDiagnosis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	if err := models.DB.Model(&models.Profile{}).Count(&totalMedicalStaff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"totalMedicine":     totalMedicine,
		"totalPatient":      totalPatient,
		"totalDiagnosis":    totalDiagnosis,
		"totalMedicalStaff": totalMedicalStaff,
	})
}
