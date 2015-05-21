package brandController

import (
	"github.com/shagtv/go-api/library/response"
	"github.com/shagtv/go-api/models"
	"net/http"
	"strconv"
)

func Info(res http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()

	brand, err := models.FindBrandByName(params.Get("name"))

	if err != nil {
		response.Error(res, err.Error(), http.StatusNotFound)
	} else {
		response.Ok(res, brand)
	}
}

func List(res http.ResponseWriter, req *http.Request) {
	skip, err := strconv.Atoi(req.URL.Query().Get("skip"))
	if err != nil {
		skip = 0
	}
	limit, err := strconv.Atoi(req.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	brands, err := models.BrandsList(skip, limit)

	if err != nil {
		response.Error(res, err.Error(), http.StatusNotFound)
	} else {
		response.Ok(res, brands)
	}
}

func Count(res http.ResponseWriter, req *http.Request) {

	count, err := models.BrandsCount()

	if err != nil {
		response.Error(res, err.Error(), http.StatusNotFound)
	} else {
		response.Ok(res, count)
	}
}
