package processor

import (
	"encoding/json"
	"github.com/Adrybe/go-driver-management-dev/internal/repository"
	"github.com/Adrybe/go-driver-management-dev/pkg/dto"
	"github.com/google/uuid"
	"log"
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

func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := repository.NewRepository()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}
	var admin dto.Admin
	err = json.NewDecoder(r.Body).Decode(&admin)
	id := uuid.New().String()
	errr := client.QueryRow(`INSERT INTO "administrators"(id,username,adminpassword,authorized)
    VALUES($1,$2,$3,$4) RETURNS id`, id, admin.UserName, admin.Password, "PENDING")
	if errr != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.Response{Description: "Admin creado"})
}
