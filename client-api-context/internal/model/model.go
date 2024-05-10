package model

type Bid struct {
	Bid string `json:"bid"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
