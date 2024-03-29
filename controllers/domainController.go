package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strings"

	"../models"
	"../utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-chi/chi"
)

var grades = []string{"F-", "F", "F+", "E-", "E", "E+", "D-", "D", "D+", "C-", "C", "C+", "B-", "B", "B+", "A-", "A", "A+"}

func InfoDomain(w http.ResponseWriter, r *http.Request) {
	name := strings.ToLower(chi.URLParam(r, "domain"))

	err := utils.CheckDomain(name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}

	host, err := Ssl(name)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	infoServer := models.InfoServer{}
	serversAux := []models.Server{}
	MinGrade := 0

	for _, endpoint := range host.Endpoints {
		domain, err := WhoIs(endpoint.IpAddress)
		if err != nil {
			continue
		}
		aux := models.Server{}
		aux.Address = endpoint.IpAddress
		aux.Grade = endpoint.Grade
		MinGrade = findMinGrade(endpoint, MinGrade)
		aux.Country = domain.CountryCode
		aux.Owner = domain.Org
		serversAux = append(serversAux, aux)
	}
	infoServer.Servers = serversAux
	infoServer.MinGrade = grades[MinGrade]

	response, err := http.Get("http://" + name)
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println("Error loading HTTP response body. ", err)
		return
	}

	infoServer.Title = document.Find("title").Text()
	infoServer.Logo = ""
	document.Find("link").Each(func(index int, element *goquery.Selection) {
		rel, exists := element.Attr("rel")
		if exists && (rel == "shortcut icon" || rel == "icon") {
			href, exists1 := element.Attr("href")
			if exists1 {
				infoServer.Logo = href
			}
		}
	})

	infoServer.PreviousMinGrade = infoServer.MinGrade
	previous := FindPreviousGrade(name)
	if previous != "" {
		infoServer.PreviousMinGrade = previous
	}

	infoServer.IsDown = false
	if host.Status == "ERROR" {
		infoServer.IsDown = true
	}
	previousServers := FindPreviousServers(name)
	infoServer.ServersChanged = !reflect.DeepEqual(previousServers, infoServer.Servers)

	Insert(name, infoServer)
	res, err := json.Marshal(infoServer)
	if err != nil {
		log.Println("Error marshaling. ", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func findMinGrade(endpoint models.Endpoint, minGrade int) int {
	if endpoint.Grade == "" {
		return minGrade
	}
	for k, v := range grades {
		if v == endpoint.Grade && (k < minGrade || minGrade == 0) {
			minGrade = k
		}
	}
	return minGrade
}
