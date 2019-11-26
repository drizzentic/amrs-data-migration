package main

import (
	"container/list"
	"fmt"
)

//What happens to programs that exist on KEMR and doesn't exist on AMRS?
var (
	HIV               = []string{"1", "10739", "dfdc6d40-2f2f-463d-ba90-cc97350441a8"}
	MCHChildServices  = []string{"2", "2050", "c2ecdf11-97cd-432a-a971-cfd9bd296b83"}
	MCHMotherServices = []string{"3", "2050", "b5d9e05f-f5ab-4612-98dd-adb75438ed34"}
	TB                = []string{"4", "10649", "9f144a34-3a4a-44a9-8486-6b7af6cc64f6"}
	IPT               = []string{"5", "10649", "335517a1-04bc-438b-9843-1ba49fb7fcd9"}
)

const updateProgram = "update amrs.program set concept_id= %s where uuid='%s'"
const insertProgram = "INSERT INTO amrs.program" +
	"( outcomes_concept_id, creator, date_created, " +
	"changed_by, date_changed, retired, name, description, uuid)" +
	" select outcomes_concept_id, creator, date_created, " +
	"changed_by, date_changed, retired, name, description, uuid " +
	"from openmrs.program where program_id=%s"

func MapProgramMetadata() {
	programList := ProgramQueueMappedValues()
	for programList.Len() > 0 {
		vals := programList.Front().Value.([]string)
		fmt.Println("-------Mapping Data Sets------")
		_, _ = amrs.Exec(fmt.Sprintf(insertProgram, vals[0]))
		_, _ = amrs.Exec(fmt.Sprintf(updateProgram, vals[1], vals[2]))
		programList.Remove(programList.Front())
	}
}
func ProgramQueueMappedValues() *list.List {
	queue := list.New()
	queue.PushBack(HIV)
	queue.PushBack(MCHChildServices)
	queue.PushBack(MCHMotherServices)
	queue.PushBack(TB)
	queue.PushBack(IPT)
	return queue
}
