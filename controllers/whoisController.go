package controllers

import (
	"encoding/json"
	"log"
	//"fmt"
	"io/ioutil"
	"net/http"

	"../models"
)

func WhoIs(domain string) (models.Domain, error) {
	response, err := http.Get("http://ip-api.com/json/" + domain)
	if err != nil {
		log.Println("The HTTP request failed with error %s\n", err)
		return models.Domain{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.Domain{}, err
	}

	ip := models.Domain{}
	err = json.Unmarshal(data, &ip)
	if err != nil {
		return models.Domain{}, err
	}

	return ip, nil
}
