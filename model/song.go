package model

import (
	"time"

	"github.com/nikitasmall/radio-go/config"
)

type Song struct {
	ID int `json:"id" sql:"AUTO_INCREMENT"`

	Singer string `json:"singer"`
	Title  string `json:"title"`

	FilePath string `json:"filePath"`

	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// init function handles migration and schema updates
func init() {
	db := config.GetConnection()
	defer db.Close()

	db.AutoMigrate(&Song{})
}
