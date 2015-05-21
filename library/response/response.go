package response
import (
	"net/http"
	"encoding/json"
)

type JsonResponse struct {
	Data interface{}
	Error string
}

func Send(res http.ResponseWriter, data interface{}, errorText string, responseStatus int) {
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

func Ok(res http.ResponseWriter, data interface{}) {
	Send(res, data, "", http.StatusOK)
}

func Error(res http.ResponseWriter, err string, responseStatus int) {
	Send(res, nil, err, responseStatus)
}