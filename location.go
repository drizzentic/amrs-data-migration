package main

var InsertLocationTag = "INSERT INTO amrs.location_tag(name, description, creator, date_created, retired, " +
	"retired_by, date_retired, retire_reason, uuid, changed_by, date_changed)" +
	" select name, description, creator, date_created, retired, " +
	"retired_by, date_retired, retire_reason, uuid, changed_by, date_changed from openmrs.location_tag"
var insertLocationAttribute = "INSERT INTO openmrs.location_attribute_type" +
	"(name, description, datatype, datatype_config, preferred_handler, handler_config," +
	"min_occurs, max_occurs, creator, date_created, changed_by, " +
	"date_changed, retired, retired_by, date_retired, " +
	"retire_reason, uuid)" +
	" select  name, description, datatype, datatype_config, preferred_handler, handler_config," +
	"min_occurs, max_occurs, creator, date_created, changed_by, " +
	"date_changed, retired, retired_by, date_retired, " +
	"retire_reason, uuid from openmrs.location_attribute_type"

//Location_tag_map
func InsertLocationMetadata() {
	amrs.Exec(InsertLocationTag)
	amrs.Exec(insertLocationAttribute)

	//Copy custom locations
	//SELECT x.* FROM amrs.location x
	//WHERE parent_location IS NOT NULL AND RETIRED= 0

	//Change uuid for standard locations
	//SELECT x.* FROM amrs.location x
	//WHERE parent_location IS NULL AND RETIRED= 0
}
