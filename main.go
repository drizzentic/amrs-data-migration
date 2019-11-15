package main

import (
	"container/list"
	"fmt"
	"log"
	"time"
)

var (
	amrs = Connect()
)

const (
	truncateTable        = "Truncate Table openmrs.%s"
	modifyConstraint     = "SET FOREIGN_KEY_CHECKS=%d;"
	truncatePersonsTable = "Truncate Table emr_migration.persons"
	insertUsers          = "Insert into openmrs.users select b.* from amrs.person a " +
		"INNER join users b on b.person_id = a.person_id;"
	insertPersons       = "Insert into openmrs.person select a.* from amrs.person a;"
	insertPersonAddress = "Insert into openmrs.person_address " +
		"(" +
		"person_address_id,person_id, preferred, address1, address2, city_village, state_province," +
		" postal_code, country, latitude, longitude, start_date, " +
		"end_date, creator, date_created, voided, voided_by, date_voided, " +
		"void_reason, county_district, address3, address4, address5, address6, " +
		"date_changed, changed_by, uuid, address7, address8, address9, " +
		"address10, address11, address12, address13, address14, address15" +
		")" +
		"select " +
		"person_address_id,person_id, preferred, address1, address2, city_village, state_province," +
		" postal_code, country, latitude, longitude, start_date, " +
		"end_date, creator, date_created, voided, voided_by, date_voided, " +
		"void_reason, county_district, address3, address4, address5, address6, " +
		"date_changed, changed_by, uuid, address7, address8, address9, " +
		"address10, address11, address12, address13, address14, address15" +
		" from amrs.person_address"
	insertPersonName = "Insert into openmrs.person_name select a.* from amrs.person_name a"
	//shift all patient attributes to kenyaemr and then update attribute type to the equivalent of kenyaemr
	insertPersonAttribute = "Insert into openmrs.person_attribute select a.* from amrs.person_attribute a"
	insertPersonAttrTypes = "Insert into openmrs.person_attribute_type(name, description, format, foreign_key, " +
		"searchable, creator, date_created, changed_by, date_changed, retired, " +
		"retired_by, date_retired, retire_reason, edit_privilege, sort_weight, uuid" +
		")" +
		"select name, description, format, foreign_key, " +
		"searchable, creator, date_created, changed_by, date_changed, retired, " +
		"retired_by, date_retired, retire_reason, edit_privilege, sort_weight, uuid " +
		"from amrs.person_attribute_type a"
	//modify to have only patients with CCC numbers

	insertPatients          = "Insert into openmrs.patient select a.* from amrs.patient a"
	insertPatientIdentifier = "Insert into openmrs.patient_identifier" +
		"(patient_identifier_id,patient_id, identifier, identifier_type, preferred, location_id, " +
		"creator, date_created, voided, voided_by, date_voided, v" +
		"oid_reason, uuid, date_changed, changed_by)" +
		"select patient_identifier_id,patient_id, identifier, identifier_type, preferred, location_id, " +
		"creator, date_created, voided, voided_by, date_voided, v" +
		"oid_reason, uuid, date_changed, changed_by from amrs.patient_identifier a"
	insertPatientIdentifierType = "Insert into openmrs.patient_identifier_type(" +
		"name, description, format, check_digit, creator, date_created, " +
		"required, format_description, validator, retired, retired_by, " +
		"date_retired, retire_reason, uuid, location_behavior, " +
		"uniqueness_behavior, date_changed, changed_by" +
		") select name, description, format, check_digit, creator, date_created, " +
		"required, format_description, validator, retired, retired_by, " +
		"date_retired, retire_reason, uuid, location_behavior, " +
		"uniqueness_behavior, date_changed, changed_by" +
		" from amrs.patient_identifier_type"
	updatePatientIdentifiers = "update openmrs.patient_identifier set identifier_type = %d where identifier_type = %d and patient_id = %d"
	updatePersonAttributes   = "update openmrs.person_attribute set person_attribute_type_id = %d where person_attribute_type_id = %d and person_id = %d"
)

func main() {
	start := time.Now()

	fmt.Println("Just getting started")
	if amrs.Ping() != nil {
		fmt.Println("Database connection to amrs could not be established")
		return
	}

	elapsed := time.Since(start)
	log.Printf("Queries execution took %s", elapsed)
	log.Println("Finished")
	defer amrs.Close()

}

func insertPerson() {
	tx, err := amrs.Begin()
	//Disable Contraint
	_, err = tx.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = tx.Exec(fmt.Sprintf(truncateTable, "person"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(fmt.Sprintf(truncateTable, "person_address"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(fmt.Sprintf(truncateTable, "person_name"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insertPersons)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insertPersonAddress)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec(insertPersonName)
	if err != nil {
		log.Fatal(err)
	}
	//Disable Contraint
	_, err = tx.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
func insertUser() {
	txu, err := amrs.Begin()
	//Disable Contraint
	_, err = txu.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		txu.Rollback()
		log.Fatal(err)
	}
	_, err = txu.Exec(fmt.Sprintf(truncateTable, "users"))
	if err != nil {
		txu.Rollback()
		log.Fatal(err)
	}
	_, err = txu.Exec(insertUsers)
	if err != nil {
		txu.Rollback()
		log.Fatal(err)
	}
	//Enable constraints
	_, err = txu.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		txu.Rollback()
		log.Fatal(err)
	}
	err = txu.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//go insertPerson()
	//insertPatient()
}
func insertPatient() {
	txi, err := amrs.Begin()
	//Disable Contraint
	_, err = txi.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		txi.Rollback()
		log.Fatal(err)
	}
	_, err = txi.Exec(fmt.Sprintf(truncateTable, "patient"))
	if err != nil {
		txi.Rollback()
		log.Fatal(err)
	}
	_, err = txi.Exec(insertPatients)
	if err != nil {
		txi.Rollback()
		log.Fatal(err)
	}
	//Enable constraints
	_, err = txi.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		txi.Rollback()
		log.Fatal(err)
	}
	err = txi.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func insertPersonAttributes() {
	txpa, err := amrs.Begin()
	//Disable Contraint
	_, err = txpa.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		txpa.Rollback()
		log.Fatal(err)
	}
	_, err = txpa.Exec(fmt.Sprintf(truncateTable, "person_attribute_type"))
	if err != nil {
		txpa.Rollback()
		log.Fatal(err)
	}
	_, err = txpa.Exec(insertPersonAttribute)
	if err != nil {
		txpa.Rollback()
		log.Fatal(err)
	}
	//Enable constraints
	_, err = txpa.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		txpa.Rollback()
		log.Fatal(err)
	}
	err = txpa.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

//Insert missing attributes types only
func insertPersonAttributeTypes() {
	txp, err := amrs.Begin()
	//Disable Contraint
	_, err = txp.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		txp.Rollback()
		log.Fatal(err)
	}
	//_, err = txp.Exec(fmt.Sprintf(truncateTable, "person_attribute_type"))
	//if err != nil {
	//	txp.Rollback()
	//	log.Fatal(err)
	//}
	_, err = txp.Exec(insertPersonAttrTypes)
	if err != nil {
		txp.Rollback()
		log.Fatal(err)
	}
	//Enable constraints
	_, err = txp.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		txp.Rollback()
		log.Fatal(err)
	}
	err = txp.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//insertPersonAttributes()
}
func insertPatientIdentifiers() {
	txpi, err := amrs.Begin()
	//Disable Contraint
	_, err = txpi.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	_, err = txpi.Exec(fmt.Sprintf(truncateTable, "patient_identifier"))
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	_, err = txpi.Exec(insertPatientIdentifier)
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	//Enable constraints
	_, err = txpi.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	err = txpi.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//insertPatientIdentifierTypes()
}

//Insert missing person attribute types only
func insertPatientIdentifierTypes() {
	txpi, err := amrs.Begin()
	//Disable Contraint
	_, err = txpi.Exec(fmt.Sprintf(modifyConstraint, 0))
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	_, err = txpi.Exec(fmt.Sprintf(truncateTable, "patient_identifier"))
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	_, err = txpi.Exec(insertPatientIdentifierType)
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	//Enable constraints
	_, err = txpi.Exec(fmt.Sprintf(modifyConstraint, 1))
	if err != nil {
		txpi.Rollback()
		log.Fatal(err)
	}
	err = txpi.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func mapPersonAttributesToAttributeTypes(a *list.Element, list *list.List) bool {
	//amrs.SetMaxOpenConns(10)

	pMap := a.Value.([]int)
	for list.Len() > 0 {
		vals := list.Front().Value.([]int)
		if len(vals) > 0 {

			amrsId := vals[0]
			kemrId := vals[1]
			if amrsId == pMap[1] {
				amrs.Exec(fmt.Sprintf(updatePersonAttributes, kemrId, pMap[1], pMap[0]))
			}
		}
		list.Remove(list.Front())
	}
	return true
}
func mapPatientIdentifiersToIdentifierTypes(a *list.Element, list *list.List) bool {
	//amrs.SetMaxOpenConns(10)

	pMap := a.Value.([]int)
	for list.Len() > 0 {
		vals := list.Front().Value.([]int)
		if len(vals) > 0 {

			amrsId := vals[0]
			kemrId := vals[1]
			if amrsId == pMap[1] {
				amrs.Exec(fmt.Sprintf(updatePatientIdentifiers, kemrId, pMap[1], pMap[0]))
				//time.Sleep(2 * time.Millisecond)
			}
		}
		list.Remove(list.Front())
	}
	return true
}

func IdMapper() {

	patientsIdQueue := CreatePatientIdentifiersQueue()
	items := IdentifiersQueueMappedValues()

	for patientsIdQueue.Len() > 0 {
		if mapPatientIdentifiersToIdentifierTypes(patientsIdQueue.Front(), items) {
			fmt.Printf("Successfully mapped %+v ", patientsIdQueue.Front().Value)
			fmt.Println("*****")
			fmt.Println("patientsIdQueue", patientsIdQueue.Len())
			patientsIdQueue.Remove(patientsIdQueue.Front())
			items = IdentifiersQueueMappedValues()
		}
	}

}
func AttributesMapper() {

	patientsAttrQueue := CreatePatientAttributesQueue()
	Attritems := AttributesQueueMappedValues()
	fmt.Println(patientsAttrQueue.Len())
	for patientsAttrQueue.Len() > 0 {
		if mapPersonAttributesToAttributeTypes(patientsAttrQueue.Front(), Attritems) {
			fmt.Printf("Successfully mapped %+v ", patientsAttrQueue.Front().Value)
			fmt.Println("*****")
			fmt.Println("patientsAttrQueue", patientsAttrQueue.Len())
			patientsAttrQueue.Remove(patientsAttrQueue.Front())
			//items = AttributesQueueMappedValues()
		}
	}

}
