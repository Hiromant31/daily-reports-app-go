package handlers

import (
	"net/http"

	"github.com/Hiromant31/daily-reports-app-go/models"

	"github.com/labstack/echo/v4"
)

func CreateReport(c echo.Context) error {
	var report models.Report
	if err := c.Bind(&report); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid data"})
	}

	// TODO: добавить в Supabase/Postgres
	// можно использовать REST API или pgx

	return c.JSON(http.StatusOK, echo.Map{"status": "received", "report": report})
}

func GetReportsByUser(c echo.Context) error {
	userID := c.Param("user_id")
	// TODO: Получить все отчёты по userID

	return c.JSON(http.StatusOK, echo.Map{
		"user_id": userID,
		"reports": []string{}, // заглушка
	})
}
