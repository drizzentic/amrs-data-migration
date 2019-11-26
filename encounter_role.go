package main

import "fmt"

var Unknown = []string{"1", "1", "a0b03050-c99b-11e0-9572-0800200c9a66"}

const mapEncounterRole = "update amrs.encounter_role set uuid = '%s' where encounter_role_id = %s"

func MapEncounterRoleMetadata() {
	amrs.Exec(fmt.Sprintf(mapEncounterRole, Unknown[2], Unknown[0]))
}
