package processor

import (
	"github.com/gorilla/mux"
	"os"
)

func SetUpApp() (string, *mux.Router) {
	port := os.Getenv("PORT")
	port = ":" + port

	r := mux.NewRouter()
	r.HandleFunc("/{admin:admin\\/?}", CreateAdmin).Methods("POST")
	return port, r
}
