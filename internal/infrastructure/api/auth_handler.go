package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/golobby/container/v3"
)

type AUTH_HANDLER func([]byte) model.LoginResponse

func handleAuthRequest(w http.ResponseWriter, r *http.Request) {
	var authUseCases contract.AuthUseCases

	if err := container.Resolve(&authUseCases); err != nil {
		log.Printf("[ERROR] Setting Up Auth Usecases: %s", err.Error())
		panic(err)
	}

	requestHandler := map[string]AUTH_HANDLER{
		"POST": authUseCases.Login,
	}[r.Method]
	body, _ := ioutil.ReadAll(r.Body)

	response := requestHandler(body)
	httpStatusCode := HTTP_CODE[response.Code]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	json.NewEncoder(w).Encode(response)
}
