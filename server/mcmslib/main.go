package mcmslib

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
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
	// API server routing
	// --------------------------------
	e := echo.New()

	e.GET("/", RootHandler) // root

	// --------------------------------
	// Start API server
	// --------------------------------
	e.Logger.Fatal(e.Start(":" + port))
}
