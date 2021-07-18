package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"strings"
)

var validate *validator.Validate

func Init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type DTO interface {
	Validate() error
}

func validateBody(r *http.Request, dto interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		fmt.Println(err)
		return err
	}

	defer r.Body.Close()

	return dto.(DTO).Validate()
}

func writeBadRequest(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}

func writeInternalServerError(w http.ResponseWriter, err error) {
	fmt.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func writeSuccess(w http.ResponseWriter, statusCode int, body interface{}) {
	response, err := json.Marshal(body)

	if err != nil {
		writeInternalServerError(w, err)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(response)
}
