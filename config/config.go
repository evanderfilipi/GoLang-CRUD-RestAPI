package config

import (
	"../structs"
	"../structs2"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dbgolang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	db.AutoMigrate(structs.Person{})
	return db
}

func DBInit2() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dbgolang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi ke database gagal!")
	}

	db.AutoMigrate(structs2.Human{})
	return db
}