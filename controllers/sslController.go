package controllers

import (
	"encoding/json"
	"log"
	//"fmt"
	"io/ioutil"
	"net/http"

	"../models"
)

func Ssl(domain string) (models.Host, error) {
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return models.Host{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.Host{}, err
	}

	host := models.Host{}
	err = json.Unmarshal(data, &host)
	if err != nil {
		return models.Host{}, err
	}

	return host, nil
}
