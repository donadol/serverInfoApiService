package controllers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

	"../models"
)

func ssl(domain string, w http.ResponseWriter) models.Host {
	response, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + domain)
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        //fmt.Println(string(data))
		ip := models.Host{}
		err = json.Unmarshal(data, &ip)
		//fmt.Printf("%+v", ip)
		return ip
	}
	return models.Host{}
}