package main

import (
	"container/list"
	"fmt"
)

// Relationship Types []int{amrs,kenyamr}
var (
	Doctor              = []string{"2", "1", "8d919b58-c2cc-11de-8d13-0010c6dffd0f"}
	Sibling             = []string{"4", "2", "8d91a01c-c2cc-11de-8d13-0010c6dffd0f"}
	Parent              = []string{"2", "3", "8d91a210-c2cc-11de-8d13-0010c6dffd0f"}
	AuntUncle           = []string{"5", "4", "8d91a3dc-c2cc-11de-8d13-0010c6dffd0f"}
	Guardian            = []string{"19", "5", "5f115f62-68b7-11e3-94ee-6bef9086de92"}
	Spouse              = []string{"7", "6", "d6895098-5d8d-11e3-94ee-b35a4132a5e3"}
	Partner             = []string{"18", "7", "007b765f-6725-4ae9-afee-9966302bace4"}
	Cowife              = []string{"10", "8", "2ac0d501-eadc-4624-b982-563c70035d46"}
	Injectabledruguser  = []string{"9"}
	Injectabledruguser2 = []string{"10"}
)

const (
	mapRelationships   = "update amrs.relationship_type set uuid = '%s' where relationship_type_id = %s"
	insertRelationship = "INSERT INTO amrs.relationship_type(a_is_to_b, b_is_to_a, preferred, weight, " +
		"description, creator, date_created, retired, retired_by, " +
		"date_retired, retire_reason, uuid) " +
		"select a_is_to_b, b_is_to_a, preferred, weight, " +
		"description, creator, date_created, retired, retired_by, " +
		"date_retired, retire_reason, uuid from openmrs.relationship_type where relationship_type_id=%s"
)

func mapRelationshipMetadata() {
	list := RelationshipQueueMappedValues()
	for list.Len() > 0 {

		vals := list.Front().Value.([]string)
		//If the attribute exists in both sides, then remap the uuid otherwise create the attribute afresh
		if len(vals) > 2 {
			kemrId := vals[0]
			uuid := vals[2]
			fmt.Println("-------Mapping Data Sets------")
			amrs.Exec(fmt.Sprintf(mapRelationships, uuid, kemrId))
		} else {
			amrs.Exec(fmt.Sprintf(insertRelationship, vals[0]))
		}
		list.Remove(list.Front())
	}
}

func RelationshipQueueMappedValues() *list.List {

	relationship := list.New()
	relationship.PushBack(Doctor)
	relationship.PushBack(Sibling)
	relationship.PushBack(Parent)
	relationship.PushBack(AuntUncle)
	relationship.PushBack(Guardian)
	relationship.PushBack(Spouse)
	relationship.PushBack(Partner)
	relationship.PushBack(Cowife)
	relationship.PushBack(Injectabledruguser)
	relationship.PushBack(Injectabledruguser2)

	return relationship
}
