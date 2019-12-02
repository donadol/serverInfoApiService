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
        log.Fatal("Error marshaling. ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}