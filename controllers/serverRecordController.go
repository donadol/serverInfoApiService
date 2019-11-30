package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"encoding/json"
	"net/http"

	"./models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	db, err = sql.Open("postgres", "postgresql://maxroach@localhost:26257/serversInfo?sslmode=disable")
	if err != nil {
		log.Fatal("Error al conectarse con la base de datos: ", err)
	}
}



