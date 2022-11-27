package utils

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestConvStringToInt1(t *testing.T) {
	r, _ := http.NewRequest("GET", "/123", nil)

	vars := map[string]string{
		"id": "123",
	}

	r = mux.SetURLVars(r, vars)

	if res, err := ConvStringToInt(r); res == 0 || err != nil {
		t.Error("test not passed")
	} else {
		t.Log("test passed")
	}
}
