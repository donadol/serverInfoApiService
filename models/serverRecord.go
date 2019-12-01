package models

type ServerRecord struct {
	Items []Item `json:"items"`
}

type Item struct {
	Domain string `json:"domain"`
}