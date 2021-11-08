package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/golobby/container/v3"
)

type AUTH_HANDLER func([]byte) (domain.LoginResponse, int)

func handleAuthRequest(w http.ResponseWriter, r *http.Request) {
	var authController domain.AuthController

	if err := container.Resolve(&authController); err != nil {
		log.Printf("[ERROR] Setting Up Auth Controller: %s", err.Error())
		panic(err)
	}

	requestHandler := map[string]AUTH_HANDLER{
		"POST": authController.Login,
	}[r.Method]

	body, _ := ioutil.ReadAll(r.Body)
	response, code := requestHandler(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}
