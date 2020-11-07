package httperr

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ErrorResponse struct {
	Error  string       `json:"error"`
	Fields []FieldError `json:"fields, omitempty"`
}

type Error struct {
	Err    error
	Code   int
	Fields []FieldError
}

func NewHttpError(err error, code int) *Error {
	return &Error{err, code, nil}
}

func (e *Error) Error() string {
	return e.Err.Error()
}
