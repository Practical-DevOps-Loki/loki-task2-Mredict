package main

import (
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"webapp/pkg/config"
)

func main() {
	port := config.GetEnv("PORT", "3000")
	logPath := config.GetEnv("LOG_PATH", "/app/log/app.log")

	e := echo.New()
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	
	if err == nil {
		multiWriter := io.MultiWriter(os.Stdout, file)
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: multiWriter,
		}))
	} else {
		e.Use(middleware.Logger())
	}

	e.Static("/", "public")

	e.GET("/", func(c echo.Context) error {
		return c.File("public/views/webapp.html")
	})

	e.Logger.Fatal(e.Start(":" + port))
}
