package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"

	"sap/trailshelper"
)


var trails = []trailshelper.Trail{}
var bikeTrails = []trailshelper.Trail{}

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

	bikeTrails = trailshelper.FilterBikeTrails(trails)

	return nil
}

func getAllTrails(c *gin.Context) {
	c.JSON(http.StatusOK, trails)
}


func getBikeTrails(c *gin.Context) {
	c.JSON(http.StatusOK, bikeTrails)
}

func getTrailsWithGrills(c *gin.Context) {
	trailsWithGrills := trailshelper.FilterTrailsWithGrills(trails)
	c.JSON(http.StatusOK, trailsWithGrills)
}

func getBikeTrailsPicnic(c *gin.Context) {
	bikeTrailsWithPicnic := trailshelper.FilterTrailsWithPicnic(bikeTrails)
	c.JSON(http.StatusOK, bikeTrailsWithPicnic)
}

func getTrailByName(c *gin.Context) {
	trail := trailshelper.FilterTrailByName(trails, c.Param("name"))
	if trail == (trailshelper.Trail{}) {
		c.JSON(http.StatusNotFound, gin.H{"message" : "trail not found"})
	} else {
		c.JSON(http.StatusOK, trail)
	}
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
	router.GET("/trails/by-name/:name", getTrailByName)
	
	router.Run("0.0.0.0:3001")
}