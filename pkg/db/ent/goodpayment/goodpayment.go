// Code generated by entc, DO NOT EDIT.

package goodpayment

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the goodpayment type in the database.
	Label = "good_payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldIdle holds the string denoting the idle field in the database.
	FieldIdle = "idle"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the goodpayment in the database.
	Table = "good_payments"
)

// Columns holds all SQL columns for goodpayment fields.
var Columns = []string{
	FieldID,
	FieldGoodID,
	FieldAccountID,
	FieldIdle,
	FieldCreateAt,
	FieldUpdateAt,
	FieldDeleteAt,
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
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() uint32
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() uint32
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() uint32
	// DefaultDeleteAt holds the default value on creation for the "delete_at" field.
	DefaultDeleteAt func() uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
