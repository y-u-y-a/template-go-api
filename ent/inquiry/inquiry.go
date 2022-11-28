// Code generated by ent, DO NOT EDIT.

package inquiry

const (
	// Label holds the string label denoting the inquiry type in the database.
	Label = "inquiry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the inquiry in the database.
	Table = "inquiries"
)

// Columns holds all SQL columns for inquiry fields.
var Columns = []string{
	FieldID,
	FieldName,
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