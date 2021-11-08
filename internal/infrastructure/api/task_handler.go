package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/ViniciusMartinsS/manager/internal/domain"
	"github.com/ViniciusMartinsS/manager/internal/domain/contract"
	"github.com/golobby/container/v3"
	"github.com/gorilla/mux"
)

type TASK_HANDLER func(domain.HandleTaskRequest) domain.TaskResponse

func handleTaskRequest(w http.ResponseWriter, r *http.Request) {
	var taskController contract.TaskController

	if err := container.Resolve(&taskController); err != nil {
		log.Printf("[ERROR] Setting Up Task Controller: %s", err.Error())
		panic(err)
	}

	requestHandler := map[string]TASK_HANDLER{
		"GET":    taskController.List,
		"POST":   taskController.Create,
		"PUT":    taskController.Update,
		"DELETE": taskController.Delete,
	}[r.Method]

	body, _ := ioutil.ReadAll(r.Body)
	userId := getUserId(r.Header)
	taskId := mux.Vars(r)["id"]

	params := domain.HandleTaskRequest{Body: body, UserId: userId, TaskId: taskId}

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
