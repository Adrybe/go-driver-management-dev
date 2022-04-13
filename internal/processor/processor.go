package processor

import (
	"net/http"
)

type (
	Processor interface {
		Test(w http.ResponseWriter, r *http.Request)
	}

	ProcessorImpl struct {
		w http.ResponseWriter
		r *http.Request
	}
)

/*func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := repository.NewRepository()
	db := client.Driver()
	col := db.Collection("barberos")

	barbero, err := validateBarbero(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{Description: "Request invalida"})
		return
	}

	_, err = col.InsertOne(context.Background(), barbero)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.Response{Description: "Barbero creado"})
}*/
