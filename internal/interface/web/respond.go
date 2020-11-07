package web

import (
	"encoding/json"
	"fmt"
	"go-ddd-api/pkg/httperr"
	"net/http"

	"github.com/pkg/errors"
)

// RespondJSON responds with json
func RespondJSON(w http.ResponseWriter, r *http.Request, data interface{}, code int) error {

	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return nil
	}

	res, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

// RespondError sends an error reponse back to the client.
func RespondError(w http.ResponseWriter, r *http.Request, err error) error {

	fmt.Println(err.Error())

	// If the error was of the type *httperr.Error, the handler has a code to use
	if webErr, ok := errors.Cause(err).(*httperr.Error); ok {
		er := httperr.ErrorResponse{
			Error:  webErr.Err.Error(),
			Fields: webErr.Fields,
		}
		if err := RespondJSON(w, r, er, webErr.Code); err != nil {
			return err
		}
		return nil
	}

	er := httperr.ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}

	if err := RespondJSON(w, r, er, http.StatusInternalServerError); err != nil {
		return err
	}

	return nil
}
