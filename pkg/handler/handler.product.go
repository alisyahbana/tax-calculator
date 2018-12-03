package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func CreateHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode("Success")
}