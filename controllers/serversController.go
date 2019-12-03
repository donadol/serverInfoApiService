package controllers

import(
	"encoding/json"
	"log"
	"net/http"

	"../models"
)

func InfoServers(w http.ResponseWriter, r *http.Request){
	record := models.ServerRecord{}
	record = FindServersRecords()
	res, err := json.Marshal(record)
	if err != nil {
		log.Println("Error marshaling. ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}