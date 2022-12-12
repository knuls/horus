package res

import (
	"net/http"

	"github.com/go-chi/render"
)

type Error struct {
	Message    string `json:"message"`
	Err        error  `json:"-"`
	StatusCode int    `json:"statusCode"`
	ErrorText  string `json:"error,omitempty"`
}

// Render implements the chi.Render interface for HTTP payload responses.
func (e *Error) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Err returns a render.Renderer generic error HTTP response.
func Err(err error, msg string, code int) render.Renderer {
	return &Error{
		Message:    msg,
		Err:        err,
		StatusCode: code,
		ErrorText:  err.Error(),
	}
}

// ErrRender returns a RenderError HTTP response.
func ErrRender(err error) render.Renderer {
	return Err(err, "response render", http.StatusUnprocessableEntity)
}

// ErrDecode returns a DecodeError HTTP response.
func ErrDecode(err error) render.Renderer {
	return Err(err, "data decode", http.StatusBadRequest)
}

// ErrNotFound returns a NotFoundError HTTP response.
func ErrNotFound(err error) render.Renderer {
	return Err(err, "not found", http.StatusNotFound)
}

// ErrBadRequest returns a BadRequestError HTTP response.
func ErrBadRequest(err error) render.Renderer {
	return Err(err, "bad request", http.StatusBadRequest)
}
