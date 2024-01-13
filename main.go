package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
	db "github.com/kkpagaev/gorofls/db/sqlc"
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

	q := db.New(conn)
	books, err := q.ListBooks(ctx, db.ListBooksParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(books)
	// Echo instance
	e := echo.New()

	// Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "foo")
}
