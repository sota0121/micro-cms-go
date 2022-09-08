package mcmslib

import (
	"fmt"
	"os"

	"github.com/labstack/echo"

	"github.com/sota0121/micro-cms-go/db"
)

func Main() {
	// ------------------------------------
	// Initialize
	// ------------------------------------
	var port = "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	fmt.Println("Start micro-cms-go server on " + port)

	// --------------------------------
	// Create DB connection
	// --------------------------------
	db := db.NewDB()
	defer db.Close()

	// --------------------------------
	// API server routing
	// --------------------------------
	e := echo.New()

	e.GET("/", RootHandler) // root

	// --------------------------------
	// Start API server
	// --------------------------------
	e.Logger.Fatal(e.Start(":" + port))
}
