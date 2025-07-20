package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	InitDB() // Подключение к базе данных

	http.HandleFunc("/sse/reports", ReportsSSEHandler)
	http.HandleFunc("/api/reports", CreateReportHandler)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
