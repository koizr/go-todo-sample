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
	server := echo.New()

	if err := godotenv.Load(); err != nil {
		server.Logger.Fatal(err)
	}

	db, err := setUpDB()
	if err != nil {
		server.Logger.Fatal(err)
	}
	handleRequest(server, &Dependencies{
		db: db,
	})
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", getPort())))
}

type ServerError struct {
	Message string `json:"message"`
}

type ServerErrorResponseBody struct {
	Error *ServerError `json:"error"`
}

func handleRequest(e *echo.Echo, dependencies *Dependencies) {
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/user", func(c echo.Context) error {
		provisionalUser := &persistent.ProvisionalUser{}
		if err := c.Bind(provisionalUser); err != nil {
			return c.JSON(http.StatusBadRequest, &struct{}{})
		}

		user, err := persistent.NewUser(provisionalUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &ServerErrorResponseBody{
				Error: &ServerError{Message: "failed to add user."},
			})
		}

		if dependencies.DB().Create(user).Error != nil {
			return c.JSON(http.StatusInternalServerError, &ServerErrorResponseBody{
				Error: &ServerError{Message: "failed to add user."},
			})
		}

		return c.JSON(http.StatusCreated, &struct{}{})
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

type Dependencies struct {
	db *gorm.DB
}

func (d *Dependencies) DB() *gorm.DB {
	return d.db
}
