package main

import (
	"fmt"
	"github.com/koizr/go-todo-sample/infra/persistent"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()

	server := echo.New()
	db, err := setUpDB()
	if err != nil {
		server.Logger.Fatal(err)
	}
	handleRequest(server, db)
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", getPort())))
}

func handleRequest(e *echo.Echo, db *gorm.DB) {
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return "80"
}

func setUpDB() (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		return
	}

	// migration
	if err = db.AutoMigrate(&persistent.User{}); err != nil {
		return
	}
	if err = db.AutoMigrate(&persistent.Task{}); err != nil {
		return
	}

	return
}
