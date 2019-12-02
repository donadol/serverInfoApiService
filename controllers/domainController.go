package controller

import(
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"

	"../models"
	"../utils"
	_ "github.com/go-chi/chi"
    _ "github.com/PuerkitoBio/goquery"
)

func InfoDomain(w http.ResponseWriter, r *http.Request){
	name := strings.ToLower(chi.URLParam(r, "domain"))
	
	err := utils.CheckDomain(name)
	if err == nil{
		log.Fatal(err)
		panic(nil)
	}

	host := Ssl(name, w)

	infoServer := models.infoServer{}
	infoServer.Servers := []models.Server{}
	MinGrade = -2

	grades := []string{"F-", "F", "F+", "E-", "E", "E+", "D-", "D", "D+", "C-", "C", "C+", "B-", "B", "B+", "A-", "A", "A+"}
	
	for _, endpoint := range host.Endpoints {
		domain := WhoIs(endpoint.IpAddress, w)
		aux := models.Server{}
		aux.Address = endpoint.IpAddress
		aux.Grade = endpoint.Grade
		index := utils.IndexOf(endpoint.Grade, grades)
		if index < MinGrade {
			grade = i
		}
		aux.Country = domain.CountryCode
		aux.Owner = domain.Org
		infoServer.Servers.append(infoServer.Servers, aux)
	}
	infoServer.MinGrade := grades[MinGrade]
	
    response, err := http.Get(name)
    if err != nil {
        log.Fatal(err)
		panic(nil)
    }
	defer response.Body.Close()
	
    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

	infoServer.Title := document.Find("title").Text()
	infoServer.Logo := ""
	document.Find("link").Each(func(index int, element *goquery.Selection) {
        rel, exists := element.Attr("rel")
        if exists && (rel == "shortcut icon" ||  rel == "icon") {
			href, exists1 := element.Attr("href")
			if exists1{
				infoServer.Logo := href
			}
        }
	})

	infoServer.PreviousMinGrade := infoServer.MinGrade
	previous := FindPreviousGrade(name)
	if previous != "" {
		infoServer.PreviousMinGrade := previous
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
        log.Fatal("Error marshaling. ", err)
	}
	w.Write(res)
}