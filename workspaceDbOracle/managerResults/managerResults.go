package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"net/http"

	"eu.corp/dbinfo"
	"eu.corp/obtainerDb"
	"github.com/gin-gonic/gin"
)

var db *sql.DB
var claimsGlb = []dbinfo.Claim{}

func main() {
	db = obtainerDb.GetDbConnection()
	claims, error := dbinfo.GetClaimsByPriority("Blocker", db)
	if error != nil {
		log.Fatal(error)
	}

	claimsGlb = claims

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

	runHttpGinServer()

}

func runHttpGinServer() {
	router := gin.Default()
	router.GET("/claims", getClaims)
	router.POST("/claims", addClaim)
	router.GET("/claim/:id", getClaimById)
	router.Run("localhost:8080")
}

func getClaims(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, claimsGlb)
}

func addClaim(c *gin.Context) {
	var newClaim dbinfo.Claim
	if err := c.BindJSON(&newClaim); err != nil {
		return
	}
	id, error := dbinfo.AddNewClaim(newClaim, db)
	if error != nil {
		log.Fatal(error)
	}
	if id > 0 {
		fmt.Println("Claim with id = %v was modified", id)
		c.IndentedJSON(http.StatusCreated, claimsGlb)
	}
}

func getClaimById(c *gin.Context) {
	id := c.Param("id")

	for _, claim := range claimsGlb {
		if strconv.Itoa(claim.ID) == id {
			c.IndentedJSON(http.StatusOK, claim)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Claim not found"})
}
