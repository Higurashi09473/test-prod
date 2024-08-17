package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	server.GET("/status", Handler)

	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	day := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(day)

	s := fmt.Sprintf("Количество дней: %d", int64(dur.Hours())/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
