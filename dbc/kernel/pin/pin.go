// Code generated by entc, DO NOT EDIT.

package pin

const (
	// Label holds the string label denoting the pin type in the database.
	Label = "pin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRid holds the string denoting the rid field in the database.
	FieldRid = "rid"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStep holds the string denoting the step field in the database.
	FieldStep = "step"
	// FieldPriority holds the string denoting the priority field in the database.
	FieldPriority = "priority"
	// FieldRelate holds the string denoting the relate field in the database.
	FieldRelate = "relate"
	// FieldUpdatedUnix holds the string denoting the updated_unix field in the database.
	FieldUpdatedUnix = "updated_unix"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// Table holds the table name of the pin in the database.
	Table = "pins"
)

// Columns holds all SQL columns for pin fields.
var Columns = []string{
	FieldID,
	FieldRid,
	FieldStatus,
	FieldStep,
	FieldPriority,
	FieldRelate,
	FieldUpdatedUnix,
	FieldComment,
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

var (
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultStep holds the default value on creation for the "step" field.
	DefaultStep string
	// DefaultPriority holds the default value on creation for the "priority" field.
	DefaultPriority int
	// DefaultRelate holds the default value on creation for the "relate" field.
	DefaultRelate string
	// DefaultUpdatedUnix holds the default value on creation for the "updated_unix" field.
	DefaultUpdatedUnix int64
)
