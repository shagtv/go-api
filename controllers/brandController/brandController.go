package brandController

import (
	"net/http"
	"encoding/json"
	"api/models"
)

type JsonResponse struct {
	Data interface{}
	Error string
}

func Info(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	brand, err := models.FindBrandByName(params.Get("name"))

	if err != nil {
		ErrorResponse(res, err.Error(), http.StatusNotFound)
	} else {
		GoodResponse(res, brand)
	}
}

func SendResponse(res http.ResponseWriter, data interface{}, errorText string, responseStatus int) {
	data = JsonResponse{Data: data, Error: errorText}

	jsonString, err := json.Marshal(data)
	if err != nil {
		data = JsonResponse{Error: err.Error()}
		responseStatus = http.StatusInternalServerError
	}

	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(responseStatus)
	res.Write(jsonString)
}

func GoodResponse(res http.ResponseWriter, data interface{}) {
	SendResponse(res, data, "", http.StatusOK)
}

func ErrorResponse(res http.ResponseWriter, err string, responseStatus int) {
	SendResponse(res, nil, err, responseStatus)
}