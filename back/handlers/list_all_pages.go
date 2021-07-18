package handlers

import (
	"back/services"
	"net/http"
)

func ListAllPages(w http.ResponseWriter, r *http.Request) {
	pages, err := services.ListAllPages()

	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	writeSuccess(w, 200, pages)
}
