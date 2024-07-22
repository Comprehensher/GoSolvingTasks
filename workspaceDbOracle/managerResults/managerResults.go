package main

import (
	"fmt"
	"log"

	"eu.corp/dbinfo"
	"eu.corp/obtainerDb"
)

func main() {
	db := obtainerDb.GetDbConnection()
	claims, error := dbinfo.GetClaimsByPriority("Blocker", db)
	if error != nil {
		log.Fatal(error)
	}
	for _, claim := range claims {
		fmt.Println(claim.CreateDate)
		fmt.Println(claim.Description)
		fmt.Println(claim.Priority)
		fmt.Println(claim.Progress)
		fmt.Println("____")
	}

	claim, err := dbinfo.GetClaimById(3, db)
	if err != nil {
		log.Fatal(err)
	}
	if claim.IsEmpty() {
		log.Println("Claim is empty")
	}
	fmt.Println("Claim is ", claim)

	fmt.Println()
	id, error := dbinfo.ModifyClaimById(3, db)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Claim id %d was updated ", id)
}
