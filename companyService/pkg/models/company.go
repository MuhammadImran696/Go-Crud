package models

import "github.com/google/uuid"

type Company struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description,string"`
	Employees   int64     `json:"employees,string,omitempty" gorm:"not null"`
	Registered  bool      `json:"registered,string,omitempty" gorm:"not null"`
	Type        string    `json:"type,string,omitempty" gorm:"not null"`
}
