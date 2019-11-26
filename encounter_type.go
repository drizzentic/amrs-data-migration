package main

import (
	"container/list"
	"fmt"
)

// Encounter Types []int{amrs,kenyamr}
//select encounter_type, year(date_created) as y from amrs.encounter group by encounter_type, y
var (
	TBScreening              = []string{"27", "1", "ed6dacc9-0827-4c82-86be-53c0d8c449be"}
	HIVDiscontinuation       = []string{"157", "2", "2bdada65-4c72-4a48-8730-859890e25cee"}
	Consultation             = []string{"167", "3", "465a92f2-baf8-42e9-9612-53064be868e8"}
	LabResults               = []string{"111", "4", "17a381d1-7e29-406a-b782-aa903b963c28"}
	Registration             = []string{"5"}
	Triage                   = []string{"110", "6", "d1059fb9-a079-4feb-a749-eedd709ae542"}
	HIVEnrollment            = []string{"7"}
	HIVConsultation          = []string{"8"}
	CWCEnrollment            = []string{"194", "9", "415f5136-ca4a-49a8-8db3-f994187c3af6"}
	CWCConsultation          = []string{"194", "10", "bcc6da85-72f2-4291-b206-789b8186a021"}
	MCHChildHEIExit          = []string{"11"}
	MCHChildImmunization     = []string{"12"}
	MCHChildDiscontinuation  = []string{"13"}
	MCHMotherEnrollment      = []string{"14"}
	MCHMotherConsultation    = []string{"15"}
	MCHMotherDiscontinuation = []string{"16"}
	TBEnrollment             = []string{"27", "17", "9d8498a4-372d-4dc4-a809-513a2434621e"}
	TBDiscontinuation        = []string{"18"}
	TBFollowUp               = []string{"19"}
	HTS                      = []string{"20"}
	ARTRefill                = []string{"21"}
	FamilyandPartnerTesting  = []string{"22"}
	HIVConfirmation          = []string{"23"}
	IPTInitiation            = []string{"25"}
	IPTOutcome               = []string{"26"}
	IPTFollowUp              = []string{"27"}
	DrugOrder                = []string{"203", "24", "7df67b83-1b84-4fe2-b1b7-794b4e9bfcc3"}
	ExternalPSmart           = []string{"28"}
	LabOrder                 = []string{"121", "30", "e1406e88-e9a9-11e8-9f32-f2801f1b9fd1"}
	DrugRegimenEditor        = []string{"29"}
	CCCDefaulterTracing      = []string{"31"}
)

const mapEncounter = "update amrs.encounter_type set uuid = '%s' where encounter_type_id = %s"
const insertEncounter = "INSERT INTO amrs.encounter_type(name, description, creator, " +
	"date_created, retired, retired_by, date_retired, " +
	"retire_reason, uuid, view_privilege, edit_privilege)" +
	"select name, description, creator, " +
	"date_created, retired, retired_by, date_retired, " +
	"retire_reason, uuid, view_privilege, edit_privilege from openmrs.encounter_type where encounter_type_id=%s"

func MapEncounterMetadata() {

	encounterList := EncounterQueueMappedValues()
	for encounterList.Len() > 0 {

		vals := encounterList.Front().Value.([]string)
		//If the attribute exists in both sides, then remap the uuid otherwise create the attribute afresh
		if len(vals) > 2 {
			kemrId := vals[0]
			uuid := vals[2]
			fmt.Println("-------Mapping Data Sets------")
			_, _ = amrs.Exec(fmt.Sprintf(mapEncounter, uuid, kemrId))
		} else {
			_, _ = amrs.Exec(fmt.Sprintf(insertEncounter, vals[0]))
		}
		encounterList.Remove(encounterList.Front())
	}

}

func EncounterQueueMappedValues() *list.List {
	queue := list.New()
	queue.PushBack(TBScreening)
	queue.PushBack(HIVDiscontinuation)
	queue.PushBack(Consultation)
	queue.PushBack(LabResults)
	queue.PushBack(Registration)
	queue.PushBack(Triage)
	queue.PushBack(HIVEnrollment)
	queue.PushBack(HIVConsultation)
	queue.PushBack(CWCEnrollment)
	queue.PushBack(CWCConsultation)
	queue.PushBack(MCHChildHEIExit)
	queue.PushBack(MCHChildImmunization)
	queue.PushBack(MCHChildDiscontinuation)
	queue.PushBack(MCHMotherEnrollment)
	queue.PushBack(MCHMotherConsultation)
	queue.PushBack(MCHMotherDiscontinuation)
	queue.PushBack(TBEnrollment)
	queue.PushBack(TBDiscontinuation)
	queue.PushBack(TBFollowUp)
	queue.PushBack(HTS)
	queue.PushBack(ARTRefill)
	queue.PushBack(FamilyandPartnerTesting)
	queue.PushBack(HIVConfirmation)
	queue.PushBack(IPTInitiation)
	queue.PushBack(IPTOutcome)
	queue.PushBack(IPTFollowUp)
	queue.PushBack(DrugOrder)
	queue.PushBack(ExternalPSmart)
	queue.PushBack(LabOrder)
	queue.PushBack(DrugRegimenEditor)
	queue.PushBack(CCCDefaulterTracing)
	//queue.PushBack(PrEPEnrollment)
	//queue.PushBack(PrEPConsultation)
	//queue.PushBack(PrEPBehaviorRiskAssessment)
	//queue.PushBack(PrEPClientDiscontinuation)
	//queue.PushBack(PrEPSTIScreening)
	//queue.PushBack(PrEPVMMCScreening)
	//queue.PushBack(FertilityIntentionScreening)
	//queue.PushBack(PrEPAllergiesscreening)
	//queue.PushBack(PrEPChronicIllnessScreening)
	//queue.PushBack(PrEPAdversedrugreactions)
	//queue.PushBack(PrEPStatus)
	//queue.PushBack(PrEPPregnancyOutcomes)
	//queue.PushBack(PrEPAppointmentcreation)
	return queue
}
