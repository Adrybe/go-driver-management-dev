package main

import (
	app "github.com/Adrybe/go-driver-management-dev/internal/processor"
	"net/http"
)

func main() {
	port, app := app.SetUpApp()
	http.ListenAndServe(port, app)
}
