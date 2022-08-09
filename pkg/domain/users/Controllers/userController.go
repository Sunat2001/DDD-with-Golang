package Controllers

import (
	UserRepository "CRUD/pkg/domain/users/Repositories"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	users := UserRepository.FetchAll()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Fatal(err)
	}
}

func GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := UserRepository.FetchById(id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user := UserRepository.FetchById(id)
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	UserRepository.Update(id, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user := UserRepository.FetchById(id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	UserRepository.Delete(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
