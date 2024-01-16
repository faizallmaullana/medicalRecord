package auth

import (
	"net/http"
	"time"

	"github.com/faizallmaullana/simrs/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InputDokter struct {
	Nip         string `json:"nip"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Specialized string `json:"specialized"`
}

// Register
// POST .../register
func Registrasi(c *gin.Context) {
	var dokter InputDokter

	if err := c.ShouldBindJSON(&dokter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the admin by username
	var user models.Users
	if err := models.DB.Where("nip = ?", dokter.Nip).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "NIP already used"})
		return
	}

	// autogenerated
	id := uuid.New().String()
	id_dokter := uuid.New().String()
	created := time.Now().UTC().Add(7 * time.Hour)

	// send data to db
	data := models.Users{
		ID:        id,
		Nip:       dokter.Nip,
		Password:  "defaultPassword",
		Role:      "dokter",
		ProfileID: id_dokter,
		CreatedAt: created,
	}

	profile := models.Profile{
		ID:          id_dokter,
		Name:        dokter.Name,
		Address:     dokter.Address,
		Specialized: dokter.Specialized,
	}

	models.DB.Create(&profile)
	models.DB.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// ===========================================================

// Registrasi Admin
// POST /api/v1/resource/admin/register
func RegisterAdmin(c *gin.Context) {
	var input GetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the admin by username
	var user models.Users
	if err := models.DB.Where("nip = ?", input.Nip).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "NIP already used"})
		return
	}

	// check length of password
	if len(input.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password length too short"})
		return
	}

	// autogenerated
	id := uuid.New().String()
	created := time.Now().UTC().Add(7 * time.Hour)

	// send data to db
	data := models.Users{
		ID:        id,
		Nip:       input.Nip,
		Password:  input.Password,
		Role:      "admin",
		CreatedAt: created,
	}

	models.DB.Create(&data)

	c.JSON(http.StatusOK, gin.H{"data": data})
}