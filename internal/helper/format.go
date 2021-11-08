package helper

import (
	"net/http"
	"strconv"
)

func GetUserId(header http.Header) int {
	headerUserId := header["User"][0]
	userId, _ := strconv.Atoi(headerUserId)
	return userId
}
