package requests

type OsSearchRequest struct {
	From        int      `json:"from"`
	Size        int      `json:"size"`
	QueryString string   `json:"queryString"`
	Index       []string `json:"index"`
}
