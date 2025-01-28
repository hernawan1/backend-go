package models

import "gorm.io/gorm"

type BadanUsaha struct {
	IdBadanUsaha   uint   `gorm:"primary key;autoIncrement" json:"id_badan_usaha"`
	NamaBadanUsaha string `json:"nama_badan_usaha"`
	NoTelp         string `json:"no_telp"`
	Email          string `json:"email"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&BadanUsaha{})
	return err
}
