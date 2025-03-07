package main

import (
	"crud-db/internal/db"
	"crud-db/internal/handler"

	"github.com/labstack/echo"
)

func main() {
	db.InitDB()

	e := echo.New()
	e.GET("/message", handler.GetHandler)
	e.POST("/message", handler.PostHandler)
	e.PATCH("message/:id", handler.PathcHandler)
	e.DELETE("message/:id", handler.DeleteHandler)
	e.Start(":8080")
}
