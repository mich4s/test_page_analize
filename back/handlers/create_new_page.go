package handlers

import (
	"back/services"
	"net/http"
)

func CreateNewPageHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	createPage := CreatePage{}

	err := validateBody(r, &createPage)

	if err != nil {
		writeBadRequest(w, err)
		return
	}

	newPage, err := services.CreateNewPage(createPage.URL)

	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	writeSuccess(w, 200, newPage)
}

type CreatePage struct {
	URL string `json:"url" validate:"required,url"`
}

func (c CreatePage) Validate() error {
	return validate.Struct(c)
}
