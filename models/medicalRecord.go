package models

import "time"

type Patient struct {
	ID      string `json:"id" gorm:"primary_key"`
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type MedicalRecord struct {
	ID string `json:"id" gorm:"primary_key"`

	// medicine
	Doses                     string `json:"doses"`
	FrequencyOfTakingMedicine string `json:"frequency_medicine"` // how much a aday
	PeriodOfTakingMedicine    string `json:"period_medicine"`    // how many days taking medicine

	// foreign key
	DokterID string  `json:"dokter_id"`
	Dokter   Profile `json:"dokter" gorm:"references:DokterID"`

	MedicineID string   `json:"medicine_id"`
	Medicine   Medicine `json:"medicine" gorm:"references:MedicineID"`

	PatientID string  `json:"patient_id"`
	Patient   Patient `json:"patient" gorm:"references:PatientID"`

	DiagnosisID string    `json:"diagnosis_id"`
	Diagnosis   Diagnosis `json:"diagnosis" gorm:"references:DiagnosisID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Diagnosis struct {
	ID        string `json:"id" gorm:"primary_key"`
	Diagnosis string `json:"diagnosis"`
	Total     int    `json:"total"`
}

type Medicine struct {
	ID       string `json:"id" gorm:"primary_key"`
	Medicine string `json:"medicine"`
	Total    int    `json:"total"`
}

type CountPatient struct {
	ID string `json:"id" gorm:"primary_key"`

	PatientID string `json:"patient_id"`
	Total     int    `json:"total"`
}
