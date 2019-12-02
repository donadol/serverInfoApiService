package models

type Item struct {
	Domain string `json:"domain"`
	Hosts []Host `json:"hosts"`
}