package personRepo

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/eotrokov/loft_api/src/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "loft"
)

func OpenConnection() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
            DB_USER, DB_PASSWORD, DB_NAME)
		db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(db)
	return db
}

func GetPerson(id uuid.UUID) *model.Person {
	var person model.Person 
    db := OpenConnection()
    err := db.QueryRow("SELECT * FROM persons WHERE id = $1", id).Scan(&person.Id, &person.LastName)
    if err != nil {
        log.Fatal(err)
	} 
	return &person
}

func GetPersons() *[]model.Person {
	db := OpenConnection()
	rows, err := db.Query("SELECT * FROM persons")
    if err != nil {
        log.Fatal(err)
	} 
	got := []model.Person{}
    for rows.Next() {
            var r model.Person
            err = rows.Scan(&r.Id, &r.LastName)
            if err != nil {
                    log.Fatalf("Scan: %v", err)
            }
            got = append(got, r)
    }
	return &got
}