// Code generated by ent, DO NOT EDIT.

package schemab

const (
	// Label holds the string label denoting the schemab type in the database.
	Label = "schemab"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the schemab in the database.
	Table = "schema_bs"
)

// Columns holds all SQL columns for schemab fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "schema_bs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"schemaa_edge_schemab",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}