package handler

// expected payload in the request
type PayloadRequest struct {
	Url    string `json:"url"`              // required
	Length string `json:"length,omitempty"` // optional
}

// payload returned
type PayloadResponse struct {
	Url string `json:"url"`
}
