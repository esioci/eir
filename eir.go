package main

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Build struct {
	ID            int
	State      	  string
	Artifacts_dir string
	Updated_at 	  string
	Inserted_at   string
	Project_id    int
}

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=esioci sslmode=disable")
	if err != nil {
		println("Error with creating connection to database.")
	}
	if err != nil {
		println("Database connection error.")
	}

	defer db.Close()

	builds := []Build{}

	// get all RUNNING builds older than 2 hours
	db.Where("(state = ? OR state = ?) AND updated_at < NOW() - Interval '2 hour'", "RUNNING", "CREATED").Find(&builds)

	for _, b := range builds {
		// change state from RUNNING to TIMEOUT
		b.State = "TIMEOUT"
		db.Save(&b)
	}
	db.Close()
}