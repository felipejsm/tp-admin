package models

type Patient struct {
	ID          uint `gorm:primarykey`
	Name        string
	Email       string
	IdTherapist string
}