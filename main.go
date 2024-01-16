package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/faizallmaullana/simrs/controller/auth"
	"github.com/faizallmaullana/simrs/controller/dokter"
	"github.com/faizallmaullana/simrs/controller/medicalrecord"
	"github.com/faizallmaullana/simrs/controller/pasien"
	"github.com/faizallmaullana/simrs/controller/statistik"
	"github.com/faizallmaullana/simrs/models"
)

// initilaize the cors middleware
var corsConfig = cors.DefaultConfig()

func init() {
	// allow all origins
	corsConfig.AllowAllOrigins = true
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// connect to database
	models.ConnectDatabase()
	r.Use(cors.New(corsConfig))

	// ROUTES
	// auth
	r.POST("/api/v1/resource/login", auth.Login)
	r.POST("/api/v1/resource/admin/register", auth.RegisterAdmin)
	r.POST("/api/v1/resource/register", auth.Registrasi)

	// dokter
	r.GET("/api/v1/resource/getdokter/:id", dokter.GetDokter)
	r.GET("/api/v1/resource/getdokter", dokter.GetDokterAll)
	r.POST("/api/v1/resource/dokter/password/change/:id", dokter.ChangePassword)

	// pasien
	r.GET("/api/v1/resource/pasien/all", pasien.GetAllPasien)

	// medical records
	r.POST("/api/v1/resource/medicalrecord/newpatient", medicalrecord.AddNewMedicalRecord)
	r.GET("/api/v1/resource/medicalrecord/getall", medicalrecord.MedicalRecordGetAll)
	r.GET("/api/v1/resource/medicalrecord/:nik", medicalrecord.MedicalRecordGetNik)
	r.GET("/api/v1/resource/medicalrecord/getBaseIdDokter/:idDokter", medicalrecord.MedicalRecordGetIdDokter)
	r.GET("/api/v1/resource/admin/get/all/obat/penyakit", medicalrecord.GetAllObatPenyakit)

	// statistics
	r.GET("/api/v1/resource/statistics", statistik.Statistics)

	// run the server
	r.Run(":9888")
}
