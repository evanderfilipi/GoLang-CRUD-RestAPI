package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	db2 := config.DBInit2()
	inDB2 := &controllers.InDB{DB: db2}

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	// router.PUT("/person", inDB.UpdatePerson)
	router.PUT("/person/:id", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)

	router.GET("/human/:id", inDB2.GetOrang)
	router.GET("/humans", inDB2.GetOrangs)
	router.POST("/human", inDB2.CreateOrang)
	// router.PUT("/human", inDB2.UpdateOrang)
	router.PUT("/human/:id", inDB2.UpdateOrang)
	router.DELETE("/human/:id", inDB2.DeleteOrang)

	router.Run(":3000")
}