package main

import (
	"fmt"
	"go_project/DB"
	"go_project/dbService"
	"go_project/pavilions"
	"go_project/shelters"
	"go_project/stream"
)

func main() {
	db, err := DB.Connect("localhost", "5432", "postgres", "22821810", "laba", "disable")
	var pavilions []pavilions.Pavilion
	pavilions, err = dbService.SelectAllPavilions(db)
	defer db.Close()
	stream.Catch(err)
	fmt.Println(pavilions)

	fmt.Println()

	var shelters []shelters.Shelter
	shelters, err = dbService.SelectAllShelters(db)
	stream.CatchF(err)
	fmt.Println(shelters)
}
