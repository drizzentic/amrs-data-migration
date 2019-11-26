package main

const insertVisitTypes = "insert into amrs.visit_type( name, description, " +
	"creator, date_created, changed_by, " +
	"date_changed, retired, retired_by, date_" +
	"retired, retire_reason, uuid)" +
	"SELECT  name, description, " +
	"creator, date_created, changed_by, " +
	"date_changed, retired, retired_by, date_" +
	"retired, retire_reason, uuid" +
	" FROM openmrs.visit_type"
const insertVisitAttr = "INSERT INTO amrs.visit_attribute_type(name, description, datatype, datatype_config," +
	" preferred_handler, handler_config, min_occurs, max_occurs, " +
	"creator, date_created, changed_by, date_changed, retired, " +
	"retired_by, date_retired, retire_reason, uuid)" +
	" select name, description, datatype, datatype_config," +
	" preferred_handler, handler_config, min_occurs, max_occurs, " +
	"creator, date_created, changed_by, date_changed, retired, " +
	"retired_by, date_retired, retire_reason, uuid from openmrs.visit_attribute_type"

func mapVisitMetadata() {

	amrs.Exec(insertVisitTypes)
	amrs.Exec(insertVisitAttr)
}
