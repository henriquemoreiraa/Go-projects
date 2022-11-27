package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ConvStringToInt(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}
