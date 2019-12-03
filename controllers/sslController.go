package controllers

import (
    "encoding/json"
	"log"
	//"fmt"
    "io/ioutil"
    "net/http"

	"../models"
)

func Ssl(domain string) models.Host {
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
        log.Fatal("The HTTP request failed with error %s\n", err)
    } else {
		data, _ := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
        //fmt.Println(string(data))
		host := models.Host{}
		err = json.Unmarshal(data, &host)
		//fmt.Printf("%+v", host)
		return host
	}
	defer response.Body.Close()
	return models.Host{}
}