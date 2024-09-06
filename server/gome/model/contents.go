package models
import (
	"gorm.io/gorm"
)

type Contents struct {
	gorm.Model
	title  string
	contents string
}
