package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/ViniciusMartinsS/manager/internal/domain/model"
	"github.com/golobby/container/v3"
	"github.com/gorilla/mux"
)

type TASK_HANDLER func(model.HandleTaskRequest) model.TaskResponse

func handleTaskRequest(w http.ResponseWriter, r *http.Request) {
	var taskUseCases contract.TaskUseCases

	if err := container.Resolve(&taskUseCases); err != nil {
		log.Printf("[ERROR] Setting Up Task Usecases: %s", err.Error())
		panic(err)
	}

	requestHandler := map[string]TASK_HANDLER{
		"GET":    taskUseCases.List,
		"POST":   taskUseCases.Create,
		"PUT":    taskUseCases.Update,
		"DELETE": taskUseCases.Delete,
	}[r.Method]

	body, _ := ioutil.ReadAll(r.Body)
	userId := getUserId(r.Header)
	taskId := mux.Vars(r)["id"]

	params := model.HandleTaskRequest{Body: body, UserId: userId, TaskId: taskId}

	response := requestHandler(params)
	httpStatusCode := HTTP_CODE[response.Code]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	json.NewEncoder(w).Encode(response)
}

func getUserId(header http.Header) int {
	headerUserId := header["User"][0]
	userId, _ := strconv.Atoi(headerUserId)
	return userId
}
