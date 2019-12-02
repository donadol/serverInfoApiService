package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"
	"net/http"

	"../models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	db, err = sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversInfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
}

func Insert(host string, info models.InfoServer) {
	stmt, err := db.Prepare("INSERT INTO DOMAIN (host, consulted_time) VALUES ($1, NOW());")
	if err != nil {
		log.Fatal("Error inserting: ", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(host)
	if err != nil {
        log.Fatal(err)
    } else{
		id, err := res.LastInsertId()
		if err != nil {
            log.Fatal(err)
        } else {
			InsertInfoServer(info, id)
        }
	}
}

func InsertInfoServer(info models.InfoServer, id int) {
	stmt, err := db.Prepare("INSERT INTO INFOSERVER (servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, domain_id) VALUES ($1, $2, $3, $4, $5, $6, $7);")
	if err != nil {
		log.Fatal("Error inserting: ", err)
	}
	defer stmt.Close()
	
	res, err := stmt.Exec(info.ServersChanged, info.MinGrade, info.PreviousMinGrade, info.Logo, info.Title, info.IsDown, id)
	if err != nil {
        log.Fatal(err)
    } else{
		id, err := res.LastInsertId()
		if err != nil {
            log.Fatal(err)
        } else {
			for _, server := range info.Servers {
				InsertServer(server, id)
			}
        }
	}
}

func InsertInfoServer(info models.Server, id int) {
	stmt, err := db.Prepare("INSERT INTO SERVER (address, ssl_grade, country, owner, infoserver_id) VALUES ($1, $2, $3, $4, $5);")
	if err != nil {
		log.Fatal("Error inserting: ", err)
	}
	defer stmt.Close()
	
	_, err := stmt.Exec(info.Address, info.Grade, info.Country, info.Owner, id)
	if err != nil {
        log.Fatal(err)
    }
}

func FindPreviousGrade(name string) (string){
	stmt, err := db.Prepare(`SELECT ssl_grade FROM DOMAIN JOIN INFOSERVER on domain_id=id
							WHERE host = $1 AND consulted_time < NOW() - INTERVAL 1 hour
							ORDER BY consulted_time DESC LIMIT 1;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	row, err := stmt.Exec(name)
	if err != nil {
        log.Fatal(err)
	}
	defer row.Close()
	var ssl
	err := row.Scan(&ssl)
	if err != nil {
		log.Fatal(err)
	}
	return ssl
}

func FindPreviousServers(name string) ([]Server]){
	stmt, err := db.Prepare(`SELECT s.address, s.ssl_grade, s.country, s.owner 
							FROM DOMAIN d, INFOSERVER is, SERVER s
							WHERE host = $1 AND d.consulted_time < NOW() - INTERVAL 1 hour 
							AND s.infoserver_id = is.id AND is.domain_id = d.id;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Exec(name)
	if err != nil {
        log.Fatal(err)
	}
	defer rows.Close()
	servers := []models.Server{}
	for rows.Next() {
		var server models.Server
        if err := rows.Scan(&server.address, &server.ssl_grade, &server.country, &server.owner); err != nil {
            log.Fatal(err)
        }
        servers = append(servers, server)
    }
	return servers
}