package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

type response struct {
	code    int
	headers map[string]string
	Data    interface{}       `json:"data,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

// Write sends the response to the client. The `response` fields can be
// overridden by passing variadic `opts` ("Functional Options") arguments. If
// no options are given, an empty `200` response is used.
func Write(w http.ResponseWriter, opts ...Option) error {
	if len(opts) == 0 {
		w.WriteHeader(http.StatusOK)
		return nil
	}

	r := &response{code: http.StatusOK}

	for _, opt := range opts {
		opt(r)
	}

	if r.code == 0 {
		return errors.New("0 is not a valid code")
	}

	for k, v := range r.headers {
		w.Header().Add(k, v)
	}

	if !isBodyAllowed(r.code) {
		w.WriteHeader(r.code)
		return nil
	}

	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.code)

	if _, err := w.Write(body); err != nil {
		return err
	}

	return nil
}

// isBodyAllowed reports whether a given response status code permits a body.
// See RFC 7230, section 3.3.
func isBodyAllowed(status int) bool {
	if (status >= 100 && status <= 199) || status == 204 || status == 304 {
		return false
	}

	return true
}

// -----------------------------------------------------------------------------

// Option helps overriding/adding response options to the current response.
type Option func(*response)

// Code sets status code.
func Code(code int) Option {
	return func(r *response) {
		r.code = code
	}
}

// Headers adds headers.
func Headers(headers map[string]string) Option {
	return func(r *response) {
		for k, v := range headers {
			r.headers[k] = v
		}
	}
}

// Success represents "successful" response.
func Success(message string, data interface{}) Option {
	return func(r *response) {
		r.Message = message
		r.Data = data
	}
}

// Error represents "failure" response.
func Error(message string, errors map[string]string) Option {
	return func(r *response) {
		r.Message = message
		r.Errors = errors
	}
}
