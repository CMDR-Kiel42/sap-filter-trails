package main

import (
	"fmt"
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

func main() {
	parseTrailCSV("BoulderTrailHeads.csv")
	router := gin.Default()
	
	router.Run("localhost:8080")
}