package main

import (
        "io"
        "log"
        "os"
        "webapp/pkg/config"

        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
)

func main() {
        port := config.GetEnv("PORT", "3000")
        logPath := config.GetEnv("LOG_PATH", "")

        e := echo.New()

        // 1. Set up file logging if LOG_PATH is provided
        var logWriter io.Writer = os.Stdout
        if logPath != "" {
                // Ensure directory exists
                err := os.MkdirAll("/app/log", 0755)
                if err != nil {
                        log.Printf("Failed to create log dir: %v", err)
                }

                // Open/create the log file
                file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
                if err == nil {
                        // MultiWriter outputs to BOTH the file and the terminal console
                        logWriter = io.MultiWriter(os.Stdout, file)
                        e.Logger.SetOutput(logWriter)
                } else {
                        log.Printf("Failed to open log file: %v", err)
                }
        }

        // 2. Attach Echo's built-in Logger middleware with your custom writer
        e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
                Output: logWriter,
        }))

        e.Static("/", "public")

        e.GET("/", func(c echo.Context) error {
                return c.File("public/views/webapp.html")
        })

        e.Start(":" + port)
}
