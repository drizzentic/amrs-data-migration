package main

import (
	"container/list"
	"log"
)

var (
	// Patient Attributes Types []int{amrs,kenyamr}
	Telephonecontact                        = []int{10, 1}
	Emailaddress                            = []int{60, 2}
	Nextofkinname                           = []int{12, 4}
	Nextofkinrelationship                   = []int{59, 5}
	Nextofkincontact                        = []int{25, 6}
	AlternatePhoneNumber                    = []int{40, 8}
	Race                                    = []int{1, 12}
	Birthplace                              = []int{2, 13}
	Citizenship                             = []int{3, 14}
	MothersName                             = []int{4, 15}
	CivilStatus                             = []int{5, 16}
	OVCHouseholdID                          = []int{6, 17}
	HealthCenter                            = []int{7, 18}
	TenCellLeader                           = []int{8, 19}
	ANCNUMBER                               = []int{9, 20}
	ContactPhoneNumber                      = []int{}
	CrisisCampLocation                      = []int{}
	NextofKin                               = []int{}
	pMTCTID1                                = []int{13, 24}
	HospitalID                              = []int{14, 25}
	HCTID1                                  = []int{15, 26}
	PointofHIVTesting                       = []int{16, 27}
	TBDistrictRegistrationNumber            = []int{17, 28}
	DistrictofBirth                         = []int{18, 29}
	HCTHouseholdID                          = []int{19, 30}
	Tribe1                                  = []int{20, 31}
	PartnerName                             = []int{21, 32}
	PartnerAge                              = []int{22, 33}
	PartnerContactPhoneNumber               = []int{23, 34}
	NextofKinAge                            = []int{24, 35}
	NextofKinContactPhoneNumber             = []int{}
	PartnerGender                           = []int{26, 37}
	NextofKinGender                         = []int{27, 38}
	TestorFakePatient                       = []int{28, 39}
	HCTLocation                             = []int{29, 40}
	HealthCenter2                           = []int{30, 41}
	Landmark                                = []int{31, 42}
	TreatmentSupporterName                  = []int{32, 43}
	TreatmentSupporterPostalAddress         = []int{33, 44}
	TreatmentSupporterContactPhoneNumber    = []int{34, 45}
	RelationshiptoTreatmentSupporter        = []int{35, 46}
	Tribe                                   = []int{36, 47}
	Commonname                              = []int{37, 48}
	Cellphoneownername                      = []int{38, 49}
	Relationshiptophonenumberowner          = []int{39, 50}
	Alternativecontactphonenumber           = []int{}
	Workplace                               = []int{41, 52}
	Occupation                              = []int{42, 53}
	Wellknownneighbourname                  = []int{43, 54}
	Plotnumber                              = []int{44, 55}
	Illnesscontactfamilymembername          = []int{45, 56}
	Landlordname                            = []int{46, 57}
	BusStage                                = []int{47, 58}
	Alternativecontactphonenumberownername  = []int{48, 59}
	ReligiousAffiliation                    = []int{49, 60}
	LocationofReligiousWorship              = []int{50, 61}
	Workplacedepartment                     = []int{51, 62}
	NameofPersonReportingInformation        = []int{52, 63}
	PhoneNumberofPersonReportingInformation = []int{53, 64}
	DrivingRoute                            = []int{54, 65}
	LocatorContactName                      = []int{55, 66}
	Relationshiptoalternativephoneowner     = []int{56, 67}
	WellKnownChildName                      = []int{57, 68}
	WellKnownSiblingName                    = []int{58, 69}
	RelationshiptoNextofKin                 = []int{}
	ContactEmailAddress                     = []int{}
	NextofKinAlternativeContactPhoneNumber  = []int{61, 72}
	BirthCompanionName                      = []int{62, 73}
	NameofReferringClinician                = []int{63, 74}
	PharmacyAnticoagulationID               = []int{64, 75}
	BloodDonorContactInformation            = []int{65, 76}
	StreetBarrack                           = []int{66, 77}
	OutreachMapPhoto                        = []int{67, 78}

	// TODO Patient Identifiers
	//The empty int arrays means the values exist in kenyaemr but are non existent in AMRS

	ACTGStudyID                 = []int{27, 41}
	AMPATHStaffPFNumber         = []int{33, 47}
	AMRSMedicalRecordNumber     = []int{3, 17}
	AMRSMedicalRecord1          = []int{35, 49}
	AMRSUniversalID             = []int{8, 22}
	AnticoagulationClinicNumber = []int{34, 48}
	CCCNumber                   = []int{28, 42}
	CCCNumber1                  = []int{29, 43}
	//CWCNumber                             = []int{}
	//DistrictRegistrationNumber            = []int{}
	//GODSNumber                            = []int{}
	HCTID                  = []int{7, 21}
	HEIIDNumber            = []int{38, 7}
	HTSNumber              = []int{}
	INVALIDCHECKDIGIT      = []int{4, 18}
	InvalidXNumberversion1 = []int{10, 24}
	KENYANNATIONALIDNUMBER = []int{5, 9}
	//KIPID                                 = []int{}
	KUZAID                = []int{40, 54}
	MTCTPlusID            = []int{2, 16}
	MTRHANCNumber         = []int{21, 35}
	MTRHAmenityNumber     = []int{17, 31}
	MTRHCARENumber        = []int{24, 38}
	MTRHCWCNumber         = []int{22, 36}
	MTRHCasualtyNumber    = []int{11, 25}
	MTRHDentalNumber      = []int{19, 33}
	MTRHENTNumber         = []int{20, 34}
	MTRHEyeClinicNumber   = []int{26, 40}
	MTRHFPNumber          = []int{13, 27}
	MTRHHospitalNumber    = []int{32, 46}
	MTRHIPNumber          = []int{15, 29}
	MTRHMemorialNumber    = []int{18, 32}
	MTRHOPDNumber         = []int{12, 26}
	MTRHORTHONumber       = []int{23, 37}
	MTRHSCCNumber         = []int{14, 28}
	MTRHStaffPFNumber     = []int{25, 39}
	MTRHXRayNumber        = []int{16, 30}
	MigrationPortVictoria = []int{31, 45}
	NHIFNumber            = []int{37, 51}
	//NationalID                            = []int{}
	//NationalUniquepatientidentifier       = []int{}
	OVCID                        = []int{36, 50}
	OldAMPATHMedicalRecordNumber = []int{1, 2}
	//OpenMRSID                             = []int{}
	//OpenMRSIdentificationNumber           = []int{}
	//PatientClinicNumber                   = []int{}
	//SmartCardSerialNumber                 = []int{}
	TemporaryRegistrationUniqueIdentifier = []int{30, 44}
	//UniquePatientNumber                   = []int{}
	XNumber      = []int{38, 23}
	ZuriHealthID = []int{39, 53}
	pMTCTID      = []int{40, 54}
)

//Push the map to a queue
func IdentifiersQueueMappedValues() *list.List {
	queue := list.New()
	queue.PushBack(pMTCTID)
	queue.PushBack(ACTGStudyID)
	queue.PushBack(AMPATHStaffPFNumber)
	queue.PushBack(AMRSMedicalRecordNumber)
	queue.PushBack(AMRSMedicalRecord1)
	queue.PushBack(AMRSUniversalID)
	queue.PushBack(AnticoagulationClinicNumber)
	queue.PushBack(CCCNumber)
	queue.PushBack(CCCNumber1)
	//queue.PushBack(CWCNumber)
	//queue.PushBack(DistrictRegistrationNumber)
	//queue.PushBack(GODSNumber)
	queue.PushBack(HCTID)
	queue.PushBack(HEIIDNumber)
	queue.PushBack(HTSNumber)
	queue.PushBack(INVALIDCHECKDIGIT)
	queue.PushBack(InvalidXNumberversion1)
	queue.PushBack(KENYANNATIONALIDNUMBER)
	//queue.PushBack(KIPID)
	queue.PushBack(KUZAID)
	queue.PushBack(MTCTPlusID)
	queue.PushBack(MTRHANCNumber)
	queue.PushBack(MTRHAmenityNumber)
	queue.PushBack(MTRHCARENumber)
	queue.PushBack(MTRHCWCNumber)
	queue.PushBack(MTRHCasualtyNumber)
	queue.PushBack(MTRHDentalNumber)
	queue.PushBack(MTRHENTNumber)
	queue.PushBack(MTRHEyeClinicNumber)
	queue.PushBack(MTRHFPNumber)
	queue.PushBack(MTRHHospitalNumber)
	queue.PushBack(MTRHIPNumber)
	queue.PushBack(MTRHMemorialNumber)
	queue.PushBack(MTRHOPDNumber)
	queue.PushBack(MTRHORTHONumber)
	queue.PushBack(MTRHSCCNumber)
	queue.PushBack(MTRHStaffPFNumber)
	queue.PushBack(MTRHXRayNumber)
	queue.PushBack(MigrationPortVictoria)
	queue.PushBack(NHIFNumber)
	//queue.PushBack(NationalID)
	//queue.PushBack(NationalUniquepatientidentifier)
	queue.PushBack(OVCID)
	queue.PushBack(OldAMPATHMedicalRecordNumber)
	//queue.PushBack(OpenMRSID)
	//queue.PushBack(OpenMRSIdentificationNumber)
	//queue.PushBack(PatientClinicNumber)
	//queue.PushBack(SmartCardSerialNumber)
	queue.PushBack(TemporaryRegistrationUniqueIdentifier)
	//queue.PushBack(UniquePatientNumber)
	queue.PushBack(XNumber)
	queue.PushBack(ZuriHealthID)
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
	queue2.PushBack(OVCHouseholdID)
	queue2.PushBack(HealthCenter)
	queue2.PushBack(TenCellLeader)
	queue2.PushBack(ANCNUMBER)
	queue2.PushBack(ContactPhoneNumber)
	queue2.PushBack(CrisisCampLocation)
	queue2.PushBack(NextofKin)
	queue2.PushBack(pMTCTID1)
	queue2.PushBack(HospitalID)
	queue2.PushBack(HCTID1)
	queue2.PushBack(PointofHIVTesting)
	queue2.PushBack(TBDistrictRegistrationNumber)
	queue2.PushBack(DistrictofBirth)
	queue2.PushBack(HCTHouseholdID)
	queue2.PushBack(Tribe1)
	queue2.PushBack(PartnerName)
	queue2.PushBack(PartnerAge)
	queue2.PushBack(PartnerContactPhoneNumber)
	queue2.PushBack(NextofKinAge)
	queue2.PushBack(NextofKinContactPhoneNumber)
	queue2.PushBack(PartnerGender)
	queue2.PushBack(NextofKinGender)
	queue2.PushBack(TestorFakePatient)
	queue2.PushBack(HCTLocation)
	queue2.PushBack(HealthCenter2)
	queue2.PushBack(Landmark)
	queue2.PushBack(TreatmentSupporterName)
	queue2.PushBack(TreatmentSupporterPostalAddress)
	queue2.PushBack(TreatmentSupporterContactPhoneNumber)
	queue2.PushBack(RelationshiptoTreatmentSupporter)
	queue2.PushBack(Tribe)
	queue2.PushBack(Commonname)
	queue2.PushBack(Cellphoneownername)
	queue2.PushBack(Relationshiptophonenumberowner)
	queue2.PushBack(Alternativecontactphonenumber)
	queue2.PushBack(Workplace)
	queue2.PushBack(Occupation)
	queue2.PushBack(Wellknownneighbourname)
	queue2.PushBack(Plotnumber)
	queue2.PushBack(Illnesscontactfamilymembername)
	queue2.PushBack(Landlordname)
	queue2.PushBack(BusStage)
	queue2.PushBack(Alternativecontactphonenumberownername)
	queue2.PushBack(ReligiousAffiliation)
	queue2.PushBack(LocationofReligiousWorship)
	queue2.PushBack(Workplacedepartment)
	queue2.PushBack(NameofPersonReportingInformation)
	queue2.PushBack(PhoneNumberofPersonReportingInformation)
	queue2.PushBack(DrivingRoute)
	queue2.PushBack(LocatorContactName)
	queue2.PushBack(Relationshiptoalternativephoneowner)
	queue2.PushBack(WellKnownChildName)
	queue2.PushBack(WellKnownSiblingName)
	queue2.PushBack(RelationshiptoNextofKin)
	queue2.PushBack(ContactEmailAddress)
	queue2.PushBack(NextofKinAlternativeContactPhoneNumber)
	queue2.PushBack(BirthCompanionName)
	queue2.PushBack(NameofReferringClinician)
	queue2.PushBack(PharmacyAnticoagulationID)
	queue2.PushBack(BloodDonorContactInformation)
	queue2.PushBack(StreetBarrack)
	queue2.PushBack(OutreachMapPhoto)

	return queue2
}

func CreatePatientIdentifiersQueue() *list.List {
	patientID := list.New()
	var (
		Patient_id int
		Identifier int
	)
	rows, err := amrs.Query("select patient_id, identifier_type from openmrs.patient_identifier order by patient_id asc limit 100")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Patient_id, &Identifier)
		if err != nil {
			log.Fatal(err)
		}
		patientID.PushBack([]int{Patient_id, Identifier})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return patientID
}
func CreatePatientAttributesQueue() *list.List {
	patientAttr := list.New()
	var (
		Person_id     int
		AttributeType int
	)
	attrows, err := amrs.Query("select person_id, person_attribute_type_id from openmrs.person_attribute order by person_id asc limit 10")
	if err != nil {
		log.Fatal(err)
	}
	defer attrows.Close()
	for attrows.Next() {
		err := attrows.Scan(&Person_id, &AttributeType)
		if err != nil {
			log.Fatal(err)
		}
		patientAttr.PushBack([]int{Person_id, AttributeType})
	}
	err = attrows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return patientAttr
}
