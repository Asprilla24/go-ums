package controllers

type Response struct {
	Error  string      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}
