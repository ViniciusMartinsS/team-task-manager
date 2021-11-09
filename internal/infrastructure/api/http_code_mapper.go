package api

import "net/http"

var HTTP_CODE = map[int]int{
	0:   http.StatusOK,
	100: http.StatusBadRequest,
	101: http.StatusUnauthorized,
	103: http.StatusForbidden,
	104: http.StatusNotFound,
	199: http.StatusInternalServerError,
}
