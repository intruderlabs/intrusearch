package responses

type GenericResponse struct {
	Success bool
	Status  int
	Body    []byte
}
