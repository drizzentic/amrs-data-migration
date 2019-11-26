package main

import (
	"container/list"
	"database/sql"
	"fmt"
	"log"
)

const checkPrivilegeExit = "select * from amrs.privilege where privilege='%s'"
const insertPrivilege = "INSERT INTO amrs.privilege" +
	"(privilege, description, uuid) " +
	"SELECT privilege, description, uuid " +
	"FROM openmrs.privilege where uuid='%s';"
const updatePrivilegeUuid = "update amrs.privilege set uuid='%s' where privilege ='%s'"

var (
	privilegeName string
	uuid          string
)

func MapPrivilegeMetadata() {
	privilege := PrivilegeQueueValues()
	for privilege.Len() > 0 {
		uuid := privilege.Front().Value.([]string)[1]
		privilegeName := privilege.Front().Value.([]string)[0]
		rows := amrs.QueryRow(fmt.Sprintf(checkPrivilegeExit, privilegeName))
		err := rows.Scan(&privilegeName, &uuid)
		if err == sql.ErrNoRows {
			amrs.Exec(fmt.Sprintf(insertPrivilege, uuid))
			fmt.Println("insert " + uuid + " ==== " + privilegeName)
		} else {
			amrs.Exec(fmt.Sprintf(updatePrivilegeUuid, uuid, privilegeName))
			fmt.Println(uuid + " ==== " + privilegeName)
		}

		privilege.Remove(privilege.Front())
	}
}
func PrivilegeQueueValues() *list.List {
	queue := list.New()

	rows, err := amrs.Query("select privilege, uuid from openmrs.privilege where privilege='App: kenyadq.dataManager'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&privilegeName, &uuid)
		if err != nil {
			log.Fatal(err)
		}
		queue.PushBack([]string{privilegeName, uuid})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return queue
}
