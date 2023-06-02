package models

import (
	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	id      int    `json:"id" gorm:"primary_key"`
	title   string `json:"title"`
	dueDate string `json:"dueDate"`
}
