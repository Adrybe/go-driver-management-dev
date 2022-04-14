package processor

import (
	"github.com/gorilla/mux"
	"os"
)

func SetUpApp() (string, *mux.Router) {
	port := os.Getenv("PORT")
	port = ":" + port

	r := mux.NewRouter()
	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/admin/{user_name}", GetAdmin).Methods("GET")
	r.HandleFunc("/admin", CreateAdmin).Methods("POST")
	return port, r
}
