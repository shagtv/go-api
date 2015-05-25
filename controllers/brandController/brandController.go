package brandController

import (
	"github.com/shagtv/go-api/library/response"
	"github.com/shagtv/go-api/storage/brandStorage"
	"net/http"
	"strconv"
)

func Info(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	brand, err := brandStorage.OneByName(params.Get("name"))

	if err != nil {
		response.Error(res, err)
	} else {
		response.Ok(res, brand)
	}
}

func List(res http.ResponseWriter, req *http.Request) {
	skip, inputError := strconv.Atoi(req.URL.Query().Get("skip"))
	if inputError != nil {
		skip = 0
	}

	limit, inputError := strconv.Atoi(req.URL.Query().Get("limit"))
	if inputError != nil || limit <= 0 {
		limit = 10
	}

	brands, err := brandStorage.List(skip, limit)

	if err != nil {
		response.Error(res, err)
	} else {
		response.Ok(res, brands)
	}
}

func Count(res http.ResponseWriter, req *http.Request) {
	count, err := brandStorage.Count()

	if err != nil {
		response.Error(res, err)
	} else {
		response.Ok(res, count)
	}
}
