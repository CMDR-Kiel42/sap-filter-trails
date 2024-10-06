package trailshelper

import "encoding/json"

type CSVBoolParsingError struct {}

func (csvError *CSVBoolParsingError) Error() string {
	return "Bad value in bool parsing"
}

type CSVBool struct {bool}

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

func FilterBikeTrails(trails []Trail) ([]Trail) {
	bikeTrails := []Trail{}
	for _, trail := range trails {
		if trail.IsBikeTrail.bool {
			bikeTrails = append(bikeTrails, trail)
		}
	}

	return bikeTrails
}

func FilterTrailsWithGrills(trails []Trail) ([]Trail) {
	trailsWithGrills := []Trail{}
	for _, trail := range trails {
		if trail.HasGrills.bool {
			trailsWithGrills = append(trailsWithGrills, trail)
		}
	}

	return trailsWithGrills
}

func FilterTrailsWithPicnic(trails []Trail) ([]Trail) {
	trailsWithPicnic := []Trail{}

	for _, trail := range trails {
		if trail.HasPicnic.bool {
			trailsWithPicnic = append(trailsWithPicnic, trail)
		}
	}

	return trailsWithPicnic
}

func FilterTrailByName(trails []Trail, name string) (Trail) {
	for _, trail := range trails {
		if trail.AccessName == name {
			return trail
		}
	}
	return Trail{}
}