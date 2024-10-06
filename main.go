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
	AccessName		string	`csv:"AccessName"`
	HasRestrooms	CSVBool	`csv:"RESTROOMS"`
	HasPicnic 		CSVBool	`csv:"PICNIC"`
	HasFishing 		CSVBool	`csv:"FISHING"`
	Aka 			string	`csv:"AKA"`
	AccessType 		string	`csv:"AccessType"`
	AccessID 		string	`csv:"AccessID"`
	Class 			string	`csv:"Class"`
	Address 		string	`csv:"Address"`
	HasFee 			CSVBool	`csv:"Fee"`
	IsBikeTrail		CSVBool	`csv:"BikeTrail"`
	HasGrills		CSVBool	`csv:"Grills"`
}

var trails = []Trail{}
var bikeTrails = []Trail{}

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

func parseTrailsCSV(fileLocation string) (err error) {
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

	bikeTrails = filterBikeTrails()

	return nil
}

func filterBikeTrails() ([]Trail) {
	bikeTrails := []Trail{}
	for _, trail := range trails {
		if trail.IsBikeTrail.bool {
			bikeTrails = append(bikeTrails, trail)
		}
	}

	return bikeTrails
}

func getAllTrails(c *gin.Context) {
	c.JSON(http.StatusOK, trails)
}


func getBikeTrails(c *gin.Context) {
	c.JSON(http.StatusOK, bikeTrails)
}

func getTrailsWithGrills(c *gin.Context) {
	trailsWithGrills := []Trail{}

	for _, trail := range trails {
		if trail.HasGrills.bool {
			trailsWithGrills = append(trailsWithGrills, trail)
		}
	}

	c.JSON(http.StatusOK, trailsWithGrills)
}

func getBikeTrailsPicnic(c *gin.Context) {
	bikeTrailsWithPicnic := []Trail{}

	for _, bikeTrail := range bikeTrails {
		if bikeTrail.HasPicnic.bool {
			bikeTrailsWithPicnic = append(bikeTrailsWithPicnic, bikeTrail)
		}
	}

	c.JSON(http.StatusOK, bikeTrailsWithPicnic)
}

func main() {
	if err := parseTrailsCSV("BoulderTrailHeads.csv"); err!= nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/trails", getAllTrails)
	router.GET("/trails/with-grills", getTrailsWithGrills)
	router.GET("/bike-trails/", getBikeTrails)
	router.GET("/bike-trails/with-picinc", getBikeTrailsPicnic)
	
	
	router.Run("localhost:8080")
}