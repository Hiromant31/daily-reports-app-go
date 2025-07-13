package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"yourmodule/handlers"
)

func main() {
	_ = godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()

	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.POST("/api/report", handlers.CreateReport)
	e.GET("/api/report/:user_id", handlers.GetReportsByUser)

	log.Println("Server listening on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
