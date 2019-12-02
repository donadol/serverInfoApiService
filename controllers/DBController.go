package controllers

import (
	"database/sql"
	"log"

	"../models"
	_ "github.com/lib/pq"
)

func Insert(host string, info models.InfoServer) {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare("INSERT INTO DOMAIN (host, consulted_time) VALUES ($1, NOW()) RETURNING id;")
	if err != nil {
		log.Fatal("Error inserting: ", err)
	}
	defer stmt.Close()
	var id int64
	err = stmt.QueryRow(host).Scan(&id)
	if err != nil {
        log.Fatal(err)
    } else{
		InsertInfoServer(info, id)
	}
	db.Close()
}

func InsertInfoServer(info models.InfoServer, id int64) {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare("INSERT INTO INFOSERVER (servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down, domain_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;")
	if err != nil {
		log.Fatal("Error inserting: ", err)
	}
	defer stmt.Close()
	var InfoServerId int64
	err = stmt.QueryRow(info.ServersChanged, info.MinGrade, info.PreviousMinGrade, info.Logo, info.Title, info.IsDown, id).Scan(&InfoServerId)
	if err != nil {
        log.Fatal(err)
    } else{
		for _, server := range info.Servers {
			InsertServer(server, InfoServerId)
		}
	}
	db.Close()
}

func InsertServer(info models.Server, id int64) {
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare("INSERT INTO SERVER (address, ssl_grade, country, owner, infoserver_id) VALUES ($1, $2, $3, $4, $5);")
	if err != nil {
		log.Fatal("Error inserting s: ", err)
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(info.Address, info.Grade, info.Country, info.Owner, id)
	if err != nil {
        log.Fatal(err)
	}
	db.Close()
}

func FindPreviousGrade(name string) (string){
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare(`SELECT infoserver.ssl_grade 
							FROM DOMAIN, INFOSERVER
							WHERE host = $1 AND consulted_time < NOW() - INTERVAL '1 hour' AND infoserver.domain_id=domain.id
							ORDER BY consulted_time DESC LIMIT 1;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	var ssl string
	err = stmt.QueryRow(name).Scan(&ssl)
	if err != nil {
        log.Print("Error in query: ",err)
		ssl=""
	}
	db.Close()
	return ssl
}

func FindPreviousServers(name string) ([]models.Server){
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare(`SELECT server.address, server.ssl_grade, server.country, server.owner 
							FROM DOMAIN, INFOSERVER, SERVER
							WHERE host = $1 AND domain.consulted_time < NOW() - INTERVAL '1 hour'
							AND server.infoserver_id = infoserver.id AND infoserver.domain_id = domain.id;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
        log.Fatal(err)
	}
	defer rows.Close()
	servers := []models.Server{}
	for rows.Next() {
		var server models.Server
		err = rows.Scan(&server.Address, &server.Grade, &server.Country, &server.Owner)
        if err != nil {
            log.Fatal(err)
        }
        servers = append(servers, server)
	}
	db.Close()
	return servers
}

func FindServersRecords()(models.ServerRecord){
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare(`SELECT DISTINCT ON (host) host, id
							FROM DOMAIN
							GROUP BY host, id
							ORDER BY host, MAX(consulted_time), id;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
        log.Fatal(err)
	}
	defer rows.Close()
	record := models.ServerRecord{}
	items := []models.Item{}

	for rows.Next() {
		var item models.Item
		var id int64
		err = rows.Scan(&item.Domain, &id)
        if err != nil {
            log.Fatal(err)
		}
		item.InfoServers = FindHosts(id)
        items = append(items, item)
	}
	record.Items = items
	db.Close()
	return record
}

func FindHosts(id int64)([]models.InfoServer){
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare(`SELECT id, servers_changed, ssl_grade, previous_ssl_grade, logo, title, is_down
							FROM INFOSERVER
							WHERE domain_id = $1;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
        log.Fatal(err)
	}
	defer rows.Close()
	hosts := []models.InfoServer{}

	for rows.Next() {
		var host models.InfoServer
		var host_id int64
		err = rows.Scan(&host_id, &host.ServersChanged, &host.MinGrade, &host.PreviousMinGrade, &host.Logo, &host.Title, &host.IsDown)
        if err != nil {
            log.Fatal(err)
		}
		host.Servers = FindServers(host_id)
        hosts = append(hosts, host)
	}
	db.Close()
	return hosts
}

func FindServers(id int64)([]models.Server){
	db, err := sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversinfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the DB: ", err)
	}
	stmt, err := db.Prepare(`SELECT address, ssl_grade, country, owner
							FROM SERVER
							WHERE infoserver_id = $1;`)
	if err != nil {
		log.Fatal("Error in select: ", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
        log.Fatal(err)
	}
	defer rows.Close()
	servers := []models.Server{}

	for rows.Next() {
		var server models.Server
		err = rows.Scan(&server.Address, &server.Grade, &server.Country, &server.Owner)
        if err != nil {
            log.Fatal(err)
		}
        servers = append(servers, server)
	}
	db.Close()
	return servers
}