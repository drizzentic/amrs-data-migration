package main

import (
	"fmt"
	"log"
	"time"
)

var (
	amrs = Connect()
)

func main() {
	start := time.Now()

	fmt.Println("Just getting started")
	if amrs.Ping() != nil {
		fmt.Println("Database connection to amrs could not be established")
		return
	}
	//Call functions to execute scriprts
	//mapVisitMetadata()
	//MapEncounterMetadata()
	//InsertLocationMetadata()
	//mapKenyaEMRPatientAttrToAMRS()
	mapKenyaEMRIdentifierTypesToAMRS()
	//mapRelationshipMetadata()
	elapsed := time.Since(start)
	log.Printf("Queries execution took %s", elapsed)
	log.Println("Finished")
	defer amrs.Close()
}
