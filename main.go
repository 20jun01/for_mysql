package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

type Country_elem struct {
	Name       string `json:"name,omitempty"  db:"Name"`
	Population int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")
	var city City
	if err := db.Get(&city, fmt.Sprintf("SELECT * FROM city WHERE Name= '%v'", os.Args[1])); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = %s", "Tokyo")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	var country Country_elem
	if err := db.Get(&country, fmt.Sprintf("SELECT Name, Population FROM country WHERE country.code = '%v'", city.CountryCode)); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = %s", "Tokyo")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	if _ ,err := db.Exec("insert into city (name, countrycode, district, population) values ('oookayama', 'JPN', 'Tokyo', 2147483647)"); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = %s", "Tokyo")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	fmt.Printf("%vの人口は%d人です\n", os.Args[1], city.Population)
	fmt.Printf("%vの人口は%vの人口の%v％です\n", os.Args[1], country.Name, float64(city.Population)/float64(country.Population) * 100)
}
