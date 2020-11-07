package web

import (
	"encoding/json"
	"go-ddd-api/pkg/httperr"
	"net/http"
	"reflect"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	validate = validator.New()
)

func init() {
	validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		// Replace struct name to json attr name
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

// DecodeJSONBody decodes and validates json body
func DecodeJSONBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return httperr.NewHttpError(err, http.StatusUnprocessableEntity)
	}

	if err := validate.Struct(v); err != nil {
		verrs, ok := err.(validator.ValidationErrors)

		if !ok {
			return err
		}

		var fields []httperr.FieldError

		for _, verr := range verrs {
			field := httperr.FieldError{
				Field: verr.Field(),
				Error: verr.Tag(),
			}

			fields = append(fields, field)
		}

		return &httperr.Error{
			Err:    err,
			Fields: fields,
			Code:   http.StatusUnprocessableEntity,
		}
	}

	return nil
}
