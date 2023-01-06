package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wwerapt/assessment/expense"

	_ "github.com/lib/pq"
)

func main() {
	//Set up Database

	//$Env:DATABASE_URL = "postgres://lfyimtvt:vGhhqviKUS8AxYYl1F3N7UmBwmXla_LH@tiny.db.elephantsql.com/lfyimtvt"
	//$Env:PORT = ":2565"

	Db := expense.GetDatabase()
	Db.InitDb()

	// Start server

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/expenses", Db.CreateExpensesHandler)
	e.GET("/expenses/:id", Db.GetIdExpensesHandler)

	log.Fatal(e.Start(os.Getenv("PORT")))

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
