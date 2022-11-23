package responses

type CreateDocumentsResponse struct {
	Total      int
	Successful int
	Failed     int
	Error      error
}
