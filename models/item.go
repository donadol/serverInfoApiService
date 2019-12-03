package models

type Item struct {
	Domain     string     `json:"domain"`
	InfoServer InfoServer `json:"infoserver"`
}
