package responses

import "io"

type GenericResponse struct {
	Success bool
	Status  int
	Body    io.ReadCloser
}
