// Code generated by ent, DO NOT EDIT.

package threat

const (
	// Label holds the string label denoting the threat type in the database.
	Label = "threat"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldThreatIDInfo holds the string denoting the threat_id_info field in the database.
	FieldThreatIDInfo = "threat_id_info"
	// FieldDomainCount holds the string denoting the domain_count field in the database.
	FieldDomainCount = "domain_count"
	// FieldTagCount holds the string denoting the tag_count field in the database.
	FieldTagCount = "tag_count"
	// FieldItelCount holds the string denoting the itel_count field in the database.
	FieldItelCount = "itel_count"
	// FieldJudge holds the string denoting the judge field in the database.
	FieldJudge = "judge"
	// FieldPoc holds the string denoting the poc field in the database.
	FieldPoc = "poc"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// Table holds the table name of the threat in the database.
	Table = "threat"
)

// Columns holds all SQL columns for threat fields.
var Columns = []string{
	FieldID,
	FieldIP,
	FieldThreatIDInfo,
	FieldDomainCount,
	FieldTagCount,
	FieldItelCount,
	FieldJudge,
	FieldPoc,
	FieldSource,
	FieldCtime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
