package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Report struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Calls     int       `json:"calls"`
}

var subscribers = make([]chan Report, 0)

func ReportsSSEHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming не поддерживается", http.StatusInternalServerError)
		return
	}

	// Настройка заголовков SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ch := make(chan Report)
	subscribers = append(subscribers, ch)

	// Отправка данных по мере поступления
	for {
		select {
		case report := <-ch:
			jsonData, _ := json.Marshal(report)
			fmt.Fprintf(w, "data: %s\n\n", jsonData)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func CreateReportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Только POST", http.StatusMethodNotAllowed)
		return
	}

	var report Report
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Ошибка JSON", http.StatusBadRequest)
		return
	}

	// Пример вставки (укажи точные поля своей таблицы reports)
	_, err = DB.Exec(`INSERT INTO reports (user_id, status, created_at, calls) VALUES ($1, $2, $3, $4)`,
		report.UserID, report.Status, time.Now(), report.Calls)
	if err != nil {
		http.Error(w, "Ошибка вставки в БД", http.StatusInternalServerError)
		log.Println("DB error:", err)
		return
	}

	// Оповестить всех подписчиков
	for _, ch := range subscribers {
		ch <- report
	}

	w.WriteHeader(http.StatusCreated)
}
