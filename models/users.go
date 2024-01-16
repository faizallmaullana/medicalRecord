package models

import "time"

type Users struct {
	ID       string `json:"id" gorm:"primary_key"`
	Nip      string `json:"nip"`
	Password string `json:"password"`
	Role     string `json:"role"`

	// foreign keys
	ProfileID string  `json:"profile_id"`
	Profile   Profile `json:"profile" gorm:"references:UserID"`

	// status
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type Profile struct {
	ID          string `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Specialized string `json:"specialized"`
}
