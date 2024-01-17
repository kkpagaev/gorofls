package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	schema "github.com/kkpagaev/gorofls/db/sqlc"
	"github.com/kkpagaev/gorofls/internal"
	"github.com/kkpagaev/gorofls/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "host=127.0.0.1 port=9666 user=user password=user dbname=user")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	e := echo.New()

	db := schema.New(conn)
	validator := &CustomValidator{validator: validator.New()}
	e.Validator = validator
	// authors := internal.NewAuthors(db)
	users := internal.NewUsers(db)
	authors := internal.NewAuthors(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/", hello)
	g := e.Group("/api")

	web.RegisterAuthorsGroup(g.Group("/authors"), web.AuthorsGroup{
		Authors: authors,
	})
	web.RegisterBookGroup(g.Group("/books"), web.BookGroup{})
	web.RegisterUserGroup(g.Group("/users"), web.UserGroup{
		Users: users,
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "foo")
}
