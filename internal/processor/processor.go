package processor

import (
	"encoding/json"
	"fmt"
	"github.com/Adrybe/go-driver-management-dev/internal/repository"
	"github.com/Adrybe/go-driver-management-dev/pkg/dto"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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

	Admin struct {
		Id         string
		UserName   string
		Password   string
		Authorized string
	}
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: home")
	fmt.Fprintf(w, "Welcome to the home page!")
}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	db, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("error geting the DB: %s", err)
	}

	var admin dto.Admin
	err = json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		log.Fatalf("error parsing request body: %+v", err)
	}

	id := uuid.New().String()

	result, err := db.Exec(`INSERT INTO public.administrators(id, username, adminpassword, authorized)
    VALUES($1, $2, $3, $4);`, id, admin.UserName, admin.Password, "PENDING")
	if err != nil {
		log.Fatalf("error inserting a new admin: %s", err)
	}

	log.Printf("%+v", result)

	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto.Response{Description: "Admin creado"})
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("error geting the DB: %s", err)
	}

	params := mux.Vars(r)

	userName := params["user_name"]

	if err != nil {
		log.Fatalf("error parsing request header: %+v", err)
	}

	result, err := db.Query(`SELECT *
		FROM public.administrators
		WHERE username = $1;`, userName)
	if err != nil {
		log.Fatalf("error fetching admin: %s", err)
	}
	var admin Admin
	for result.Next() {
		var id, username, adminpassword, authorized string
		result.Scan(&id, &username, &adminpassword, &authorized)
		admin = Admin{
			Id:         id,
			UserName:   username,
			Password:   adminpassword,
			Authorized: authorized,
		}
		log.Printf("el administrador encontrado es %+v", admin)

	}

	log.Printf("%+v", result)

	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

func SignInAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := repository.NewRepository()
	if err != nil {
		log.Fatalf("error geting the DB: %s", err)
	}

	var request dto.Admin
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		log.Fatalf("error parsing request body: %+v", err)
	}

	if err != nil {
		log.Fatalf("error parsing request header: %+v", err)
	}

	result, err := db.Query(`SELECT *
		FROM public.administrators
		WHERE username = $1;`, request.UserName)
	if err != nil {
		log.Fatalf("error fetching admin: %s", err)
	}
	var admin Admin
	for result.Next() {
		var id, username, adminpassword, authorized string
		result.Scan(&id, &username, &adminpassword, &authorized)
		admin = Admin{
			Id:         id,
			UserName:   username,
			Password:   adminpassword,
			Authorized: authorized,
		}
		log.Printf("el administrador encontrado es %+v", admin)
	}

	log.Printf("%+v", result)
	defer db.Close()
	if passwordVerifier(admin, request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(admin)
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(dto.Response{Description: "Error en la password."})
}

func passwordVerifier(dbUser Admin, requestUser dto.Admin) bool {
	return dbUser.Password == requestUser.Password
}
