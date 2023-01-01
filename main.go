package main

import (
	"github.com/go-verify/pkg/controllers/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/verify", http.Verify)
	e.Logger.Fatal(e.Start(":" + port))
}
