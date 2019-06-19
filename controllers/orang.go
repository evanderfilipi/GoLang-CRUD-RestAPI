package controllers

import (
	"net/http"

	"../structs2"
	"github.com/gin-gonic/gin"
)

// to get one data with {id}
func (idb *InDB) GetOrang(c *gin.Context) {
	var (
		human structs2.Human
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&human).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": human,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in person
func (idb *InDB) GetOrangs(c *gin.Context) {
	var (
		humans []structs2.Human
		result  gin.H
	)

	idb.DB.Find(&humans)
	if len(humans) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": humans,
			"count":  len(humans),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreateOrang(c *gin.Context) {
	var (
		orang structs2.Human
		result gin.H
	)
	nama_awal := c.PostForm("nama_awal")
	nama_akhir := c.PostForm("nama_akhir")
	no_telepon := c.PostForm("no_telepon")
	umur := c.PostForm("umur")
	kota := c.PostForm("kota")
	orang.Nama_Awal = nama_awal
	orang.Nama_Akhir = nama_akhir
	orang.No_Telepon = no_telepon
	orang.Umur = umur
	orang.Kota = kota
	idb.DB.Create(&orang)
	result = gin.H{
		"result": orang,
	}
	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (idb *InDB) UpdateOrang(c *gin.Context) {
	// id := c.Query("id")
	id := c.Param("id")
	nama_awal := c.PostForm("nama_awal")
	nama_akhir := c.PostForm("nama_akhir")
	no_telepon := c.PostForm("no_telepon")
	umur := c.PostForm("umur")
	kota := c.PostForm("kota")
	var (
		orang    structs2.Human
		newOrang structs2.Human
		result    gin.H
	)

	err := idb.DB.Where("id = ?", id).First(&orang).Error
	if err != nil {
		result = gin.H{
			"result": "Data tidak ditemukan!",
		}
	} else {
		newOrang.Nama_Awal = nama_awal
		newOrang.Nama_Akhir = nama_akhir
		newOrang.No_Telepon = no_telepon
		newOrang.Umur = umur
		newOrang.Kota = kota
		err = idb.DB.Model(&orang).Updates(newOrang).Error
		if err != nil {
			result = gin.H{
				"result": "Data gagal diupdate!",
			}
		} else {
			result = gin.H{
				"result": "Data berhasil diupdate!",
			}
		}
	}
	
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteOrang(c *gin.Context) {
	var (
		orang structs2.Human
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&orang).Error
	if err != nil {
		result = gin.H{
			"result": "Data tidak ditemukan!",
		}
	} else {
		err = idb.DB.Delete(&orang).Error
		if err != nil {
			result = gin.H{
				"result": "Data gagal dihapus!",
			}
		} else {
			result = gin.H{
				"result": "Data berhasil dihapus!",
			}
		}
	}

	c.JSON(http.StatusOK, result)
}
