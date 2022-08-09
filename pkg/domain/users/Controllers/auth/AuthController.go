package auth

import (
	"CRUD/pkg/config"
	"CRUD/pkg/domain/users/Entities"
	UserRepository "CRUD/pkg/domain/users/Repositories"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type AppSecret struct {
	Key string
}

type loginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	user := Entities.User{}
	user = UserRepository.FetchByEmail(loginStruct.Email, user)

	if user.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not fount with email:" + loginStruct.Email)
		return
	}

	isPasswordCorrect := user.CheckPassword(loginStruct.Password)

	if !isPasswordCorrect {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Incorrect Password")
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	var SECRETtype AppSecret

	err = envconfig.Process("APP", &SECRETtype)
	if err != nil {
		log.Fatal(err)
	}

	var SECRET = []byte(SECRETtype.Key)

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenStr)
	return

}

func validateLogin(w http.ResponseWriter, r *http.Request) (*loginPayload, error) {
	var loginStruct loginPayload

	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = config.ValidateStruct(loginStruct)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user Entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = config.ValidateStruct(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	user = UserRepository.Create(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	user.HidePassword()
	// TODO think about return user or Bearer token
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}
