package medicalrecord

import (
	"net/http"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
)

// ambil semua data medicalrecord
// GET /medicalrecord/getall
func MedicalRecordGetAll(c *gin.Context) {
	var Pasien []models.MedicalRecord
	if err := models.DB.Preload("Patient").Preload("Medicine").Preload("Diagnosis").Preload("Dokter").Find(&Pasien).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access this"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Pasien})
}

// ambil data medical record berdasar nik
// GET /medicalrecord/:nik
func MedicalRecordGetNik(c *gin.Context) {
	var Pasien models.Patient
	if err := models.DB.Where("nik = ?", c.Param("nik")).First(&Pasien).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datamu Tidak Ditemukan"})
		return
	}

	var Data []models.MedicalRecord
	if err := models.DB.Where("patient_id = ?", Pasien.ID).Preload("Medicine").Preload("Diagnosis").Preload("Dokter").Preload("Patient").Find(&Data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datamu Tidak Ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Data})
}

// GET /medicalrecord/getBaseIdDokter/:idDokter
func MedicalRecordGetIdDokter(c *gin.Context) {
	var Dokter models.Users
	if err := models.DB.Where("id = ?", c.Param("idDokter")).First(&Dokter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datamu Tidak Ditemukan"})
		return
	}

	var medicalRecord []models.MedicalRecord
	if err := models.DB.Where("dokter_id = ?", Dokter.ProfileID).Preload("Patient").Preload("Medicine").Preload("Diagnosis").Preload("Dokter").Find(&medicalRecord).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access this"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicalRecord})
}
