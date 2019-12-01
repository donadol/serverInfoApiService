package controllers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

	"../models"
)

func WhoIs(domain string) models.Domain {
	response, err := http.Get("http://ip-api.com/json/" + domain)
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        //fmt.Println(string(data))
		ip := models.Domain{}
		err = json.Unmarshal(data, &ip)
		//fmt.Printf("%+v", ip)
		return ip
	}
	return models.Domain{}
}