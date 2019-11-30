package models

type Server struct {
	Address string `json:"address"`
	Grade string `json:"ssl_grade"`
	Country string `json:"country"`
	Owner string `json:"owner"`
}