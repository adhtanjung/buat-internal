package models

type WebResponse struct {
	Status int32       `json:"status"`
	Data   interface{} `json:"data"`
}

type RegisterResponse struct {
	Status  int32               `json:"status"`
	Error   []map[string]string `json:"errors"`
	Message string              `json:"message"`
}

type LoginResponse struct {
	Status int32  `json:"status"`
	Token  string `json:"token"`
}
