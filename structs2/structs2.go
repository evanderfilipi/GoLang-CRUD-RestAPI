package structs2

import "github.com/jinzhu/gorm"

type Human struct {
  gorm.Model
  Nama_Awal string
  Nama_Akhir string
  No_Telepon string
  Umur string
  Kota string
}

func (Human) TableName() string {
    return "Orang"
}