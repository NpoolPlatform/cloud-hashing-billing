// Code generated by ent, DO NOT EDIT.

package userpaymentbalance

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the userpaymentbalance type in the database.
	Label = "user_payment_balance"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldPaymentID holds the string denoting the payment_id field in the database.
	FieldPaymentID = "payment_id"
	// FieldUsedByPaymentID holds the string denoting the used_by_payment_id field in the database.
	FieldUsedByPaymentID = "used_by_payment_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCoinUsdCurrency holds the string denoting the coin_usd_currency field in the database.
	FieldCoinUsdCurrency = "coin_usd_currency"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// Table holds the table name of the userpaymentbalance in the database.
	Table = "user_payment_balances"
)

// Columns holds all SQL columns for userpaymentbalance fields.
var Columns = []string{
	FieldID,
	FieldAppID,
	FieldUserID,
	FieldPaymentID,
	FieldUsedByPaymentID,
	FieldAmount,
	FieldCoinUsdCurrency,
	FieldCoinTypeID,
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
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
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
