package models

type ResponseData struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
