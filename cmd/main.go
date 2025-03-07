package main

import (
	"C:\Users\Rodion\stepic\internal\db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	initDB()
	e := echo.New()
	e.GET("/message", GetHandler)
	e.POST("/message", PostHandler)
	e.PATCH("message/:id", PathcHandler)
	e.DELETE("message/:id", DeleteHandler)
	e.Start(":8080")
}
