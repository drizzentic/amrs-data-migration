package main

import (
	"container/list"
	"fmt"
)

var (
	// Patient Attributes Types []int{amrs,kenyamr}
	Telephonecontact      = []string{"10", "8", "b2c38640-2603-4629-aebd-3b54f33f1e3a"}
	Emailaddress          = []string{"60", "9", "b8d0b331-1d2d-4a9a-b741-1816f498bdb6"}
	Nextofkinname         = []string{"12", "11", "830bef6d-b01f-449d-9f8d-ac0fede8dbd3"}
	Nextofkinrelationship = []string{"59", "12", "d0aa9fd1-2ac5-45d8-9c5e-4317c622c8f5"}
	Nextofkincontact      = []string{"25", "13", "342a1d39-c541-4b29-8818-930916f4c2dc"}
	AlternatePhoneNumber  = []string{"40", "15", "94614350-84c8-41e0-ac29-86bc107069be"}
	Race                  = []string{"1", "1", "8d871386-c2cc-11de-8d13-0010c6dffd0f"}
	Birthplace            = []string{"2", "2", "8d8718c2-c2cc-11de-8d13-0010c6dffd0f"}
	Citizenship           = []string{"3", "3", "8d871afc-c2cc-11de-8d13-0010c6dffd0f"}
	MothersName           = []string{"4", "4", "8d871d18-c2cc-11de-8d13-0010c6dffd0f"}
	CivilStatus           = []string{"5", "5", "8d871f2a-c2cc-11de-8d13-0010c6dffd0f"}
	NextOfKinAddress      = []string{"14§"}
	SubchiefName          = []string{"10"}
	GuardianFirstName     = []string{"17"}
	GuardianLastName      = []string{"18"}
	HealthCenter          = []string{"7", "7", "8d87236c-c2cc-11de-8d13-0010c6dffd0f"}
	HealthDistrict        = []string{"6"}
	NearestHealthFacility = []string{"16"}
	// TODO Patient Identifiers
	//The empty int arrays means the values exist in kenyaemr but are non existent in AMRS
	CWCNumber                       = []string{"22", "10", "1dc8b419-35f2-4316-8d68-135f0689859b"}
	DistrictRegistrationNumber      = []string{"8"}
	GODSNumber                      = []string{"13"}
	HEIIDNumber                     = []string{"38", "7", "0691f522-dd67-4eeb-92c8-af5083baf338"}
	HTSNumber                       = []string{"12"}
	KIPID                           = []string{"14"}
	NHIFNumber                      = []string{"37", "15", "09ebf4f9-b673-4d97-b39b-04f94088ba64"}
	NationalID                      = []string{"5", "5", "49af6cdc-7968-4abb-bf46-de10d7f4859f"}
	NationalUniquepatientidentifier = []string{"9"}
	OpenMRSID                       = []string{"3"}
	OpenMRSIdentificationNumber     = []string{"3", "1", "8d793bee-c2cc-11de-8d13-0010c6dffd0f"}
	PrepUniqueNumber                = []string{"16"}
	PatientClinicNumber             = []string{"4"}
	SmartCardSerialNumber           = []string{"11"}
	UniquePatientNumber             = []string{"28", "6", "f2d6ff1a-8440-4d35-a150-1d4b5a930c5e"}
)

const (
	insertPersonAttrTypes = "Insert into amrs.person_attribute_type(name, description, format, foreign_key, " +
		"searchable, creator, date_created, changed_by, date_changed, retired, " +
		"retired_by, date_retired, retire_reason, edit_privilege, sort_weight, uuid" +
		")" +
		"select name, description, format, foreign_key, " +
		"searchable, creator, date_created, changed_by, date_changed, retired, " +
		"retired_by, date_retired, retire_reason, edit_privilege, sort_weight, uuid " +
		"from openmrs.person_attribute_type a where person_attribute_type_id=%s"
	insertPatientIdentifierType = "Insert into amrs.patient_identifier_type(" +
		"name, description, format, check_digit, creator, date_created, " +
		"required, format_description, validator, retired, retired_by, " +
		"date_retired, retire_reason, uuid, location_behavior, " +
		"uniqueness_behavior, date_changed, changed_by" +
		") select name, description, format, check_digit, creator, date_created, " +
		"required, format_description, validator, retired, retired_by, " +
		"date_retired, retire_reason, uuid, location_behavior, " +
		"uniqueness_behavior, date_changed, changed_by" +
		" from openmrs.patient_identifier_type where patient_identifier_type_id=%s"
	updatePatientIdentifiers  = "update amrs.patient_identifier_type set uuid = %s where patient_identifier_type_id = %s"
	updatePersonAttributeType = "update amrs.person_attribute_type set uuid = '%s' where person_attribute_type_id =%s"
)

//Push the map to a queue
func IdentifiersQueueMappedValues() *list.List {
	queue := list.New()

	queue.PushBack(CWCNumber)
	queue.PushBack(DistrictRegistrationNumber)
	queue.PushBack(GODSNumber)
	queue.PushBack(HEIIDNumber)
	queue.PushBack(HTSNumber)
	queue.PushBack(KIPID)
	queue.PushBack(NHIFNumber)
	queue.PushBack(NationalID)
	queue.PushBack(NationalUniquepatientidentifier)
	queue.PushBack(OpenMRSID)
	queue.PushBack(OpenMRSIdentificationNumber)
	queue.PushBack(PatientClinicNumber)
	queue.PushBack(SmartCardSerialNumber)
	queue.PushBack(UniquePatientNumber)
	queue.PushBack(PrepUniqueNumber)

	return queue
}

//Push the map to a queue
func AttributesQueueMappedValues() *list.List {
	queue2 := list.New()
	queue2.PushBack(Telephonecontact)
	queue2.PushBack(Emailaddress)
	queue2.PushBack(Nextofkinname)
	queue2.PushBack(Nextofkinrelationship)
	queue2.PushBack(Nextofkincontact)
	queue2.PushBack(AlternatePhoneNumber)
	queue2.PushBack(Race)
	queue2.PushBack(Birthplace)
	queue2.PushBack(Citizenship)
	queue2.PushBack(MothersName)
	queue2.PushBack(CivilStatus)
	queue2.PushBack(HealthCenter)
	queue2.PushBack(NextOfKinAddress)
	queue2.PushBack(SubchiefName)
	queue2.PushBack(GuardianFirstName)
	queue2.PushBack(GuardianLastName)
	queue2.PushBack(HealthDistrict)
	queue2.PushBack(NearestHealthFacility)
	return queue2
}
func mapKenyaEMRPatientAttrToAMRS() {
	list := AttributesQueueMappedValues()
	for list.Len() > 0 {

		vals := list.Front().Value.([]string)
		//If the attribute exists in both sides, then remap the uuid otherwise create the attribute afresh
		if len(vals) > 2 {
			kemrId := vals[0]
			uuid := vals[2]
			fmt.Println("-------Mapping Data Sets------")
			amrs.Exec(fmt.Sprintf(updatePersonAttributeType, uuid, kemrId))
		} else {
			amrs.Exec(fmt.Sprintf(insertPersonAttrTypes, vals[0]))
		}
		list.Remove(list.Front())
	}
}
func mapKenyaEMRIdentifierTypesToAMRS() {
	list := IdentifiersQueueMappedValues()
	for list.Len() > 0 {

		vals := list.Front().Value.([]string)
		//If the attribute exists in both sides, then remap the uuid otherwise create the attribute afresh
		if len(vals) > 2 {
			kemrId := vals[0]
			uuid := vals[2]
			fmt.Println("-------Mapping Data Sets------")
			amrs.Exec(fmt.Sprintf(updatePatientIdentifiers, uuid, kemrId))
		} else {
			amrs.Exec(fmt.Sprintf(insertPatientIdentifierType, vals[0]))
		}
		list.Remove(list.Front())
	}
}
