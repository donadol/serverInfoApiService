package models

type Item struct {
	Domain string `json:"domain"`
	InfoServers []InfoServer `json:"infoservers"`
}