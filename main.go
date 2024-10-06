package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
)


type CSVBoolParsingError struct {}

func (csvError *CSVBoolParsingError) Error() string {
	return "Bad value in bool parsing"
}

type CSVBool struct {
	bool
}

type Trail struct {
	FID 			int		`csv:"FID"`
	HasRestrooms	CSVBool	`csv:"RESTROOMS"`
	HasPicnic 		CSVBool	`csv:"PICNIC"`
	HasFishing 		CSVBool	`csv:"FISHING"`
	Aka 			string	`csv:"AKA"`
	AccessType 		string	`csv:"AccessType"`
	AccessID 		string	`csv:"AccessID"`
	Class 			string	`csv:"Class"`
	Address 		string	`csv:"Address"`
	HasFee 			string	`csv:"Fee"`
	IsBikeTrail		CSVBool	`csv:"BikeTrail"`
}

var trails = []Trail{}


func (csvBool *CSVBool) UnmarshalCSV(csvValue string) (err error) {
	if csvValue == "Yes" {
		csvBool.bool, err = true, nil
	} else if csvValue == "No" {
		csvBool.bool, err = false, nil
	} else {
		return &CSVBoolParsingError{}
	}

	return err
}

func (csvBool *CSVBool) MarshalJSON() ([]byte, error) {
	parsedValue := csvBool.bool

	return json.Marshal(parsedValue)
}

func parseTrailCSV(fileLocation string) (err error) {
	csvFile, err := os.Open(fileLocation)
	if err != nil {
		fmt.Printf("couldn't open csv at location %v", fileLocation)
		return err
	}
	defer csvFile.Close()

	if err := gocsv.UnmarshalFile(csvFile, &trails); err != nil {
		fmt.Printf("couldn't parse csv at location %v", fileLocation)
		return err
	}

	return nil
}

func getAllTrails(c *gin.Context) {
	c.JSON(http.StatusOK, trails)
}

func getBikeTrails(c *gin.Context) {
	bikeTrails := []Trail{}

	for _, trail := range trails {
		if trail.IsBikeTrail.bool {
			bikeTrails = append(bikeTrails, trail)
		}
	}

	c.JSON(http.StatusOK, bikeTrails)
}

func main() {
	if err := parseTrailCSV("BoulderTrailHeads.csv"); err!= nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/trails", getAllTrails)
	router.GET("/bike-trails/", getBikeTrails)
	
	router.Run("localhost:8080")
}