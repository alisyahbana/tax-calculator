package handler

import (
	"encoding/json"
	"github.com/alisyahbana/tax-calculator/pkg/common/log"
	"github.com/alisyahbana/tax-calculator/pkg/service/product"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"message"`
}

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jsonBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response := ErrorResponse{
			Error: err.Error(),
		}
		log.Error(err.Error())
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(response)
		return
	}

	var productInputs []product.ProductInput

	err = json.Unmarshal(jsonBody, &productInputs)
	if err != nil {
		response := ErrorResponse{
			Error: err.Error(),
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(response)
		return
	}

	var newProductIds []uint64

	for _, productInput := range productInputs {
		newProductId, err := product.New().CreateProduct(productInput)
		if err != nil {
			log.Error(err.Error())
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(err.Error())
			return
		}

		newProductIds = append(newProductIds, newProductId)
	}

	billing, err := product.New().GenerateBilling(productInputs)
	if err != nil {
		log.Error(err.Error())
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err.Error())
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(billing)

}
