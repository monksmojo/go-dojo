package models

import (
	"github.com/monks_mojo/go-dojo/go-projects/go-bookstore-managment-asyt/pkg/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Author      string `json:"author" gorm:"not null"`
	Publication string `json:"publication" gorm:"not null"`
}

// gorm.Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt It may be embedded into your model or you may build your own model without it



func (book *Book) CreateBook() *Book {
	db.NewRecord()
}
