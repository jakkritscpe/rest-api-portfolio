package models

import "gorm.io/gorm"

type Tools struct {
	Name     string
	Urlimg   string
	Category string
	gorm.Model
}

type Tools_Category struct {
	Name string
	gorm.Model
}
