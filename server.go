package main

import (
	"errors"
	"fmt"
	"github.com/koizr/go-todo-sample/auth/app"
	"github.com/koizr/go-todo-sample/common"
	"github.com/koizr/go-todo-sample/infra/persistent"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"

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
		return
	}
	secret, err := getSecret()
	if err != nil {
		server.Logger.Fatal(err)
		return
	}
	handleRequest(server, &Dependencies{
		db:     db,
		secret: secret,
	})
	server.Logger.Fatal(server.Start(fmt.Sprintf(":%s", getPort())))
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
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to add user."))
		}

		if dependencies.DB().Create(user).Error != nil {
			return c.JSON(http.StatusInternalServerError, common.NewServerError("failed to add user."))
		}

		return c.JSON(http.StatusCreated, &struct{}{})
	})
	e.POST("/login", app.Login(dependencies))

	_ = middleware.JWT(dependencies.Secret())
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return "80"
}

func getSecret() (string, error) {
	if secret := os.Getenv("SECRET_KEY"); secret != "" {
		return secret, nil
	} else {
		return "", errors.New("secret key not found")
	}
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
	db     *gorm.DB
	secret string
}

func (d *Dependencies) DB() *gorm.DB {
	return d.db
}

func (d *Dependencies) Secret() string {
	return d.secret
}

func (d *Dependencies) Now() *time.Time {
	now := time.Now()
	return &now
}
