package main

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	schema "github.com/kkpagaev/gorofls/db/sqlc"
	"github.com/kkpagaev/gorofls/internal"
	"github.com/kkpagaev/gorofls/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Id int `json:"id"`
}

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "host=127.0.0.1 port=9666 user=user password=user dbname=user")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	db := schema.New(conn)
	// authors := internal.NewAuthors(db)
	users := internal.NewUsers(db)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", hello)
	g := e.Group("/api")

	web.BookGroup(g.Group("/books"))
	web.UserGroup(g.Group("/users"), web.UserGroupDeps{
		Users: users,
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "foo")
}
