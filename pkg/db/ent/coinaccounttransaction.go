// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccounttransaction"
	"github.com/google/uuid"
)

// CoinAccountTransaction is the model entity for the CoinAccountTransaction schema.
type CoinAccountTransaction struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// FromAddressID holds the value of the "from_address_id" field.
	FromAddressID uuid.UUID `json:"from_address_id,omitempty"`
	// ToAddressID holds the value of the "to_address_id" field.
	ToAddressID uuid.UUID `json:"to_address_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount uint64 `json:"amount,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// State holds the value of the "state" field.
	State coinaccounttransaction.State `json:"state,omitempty"`
	// ChainTransactionID holds the value of the "chain_transaction_id" field.
	ChainTransactionID string `json:"chain_transaction_id,omitempty"`
	// PlatformTransactionID holds the value of the "platform_transaction_id" field.
	PlatformTransactionID uuid.UUID `json:"platform_transaction_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinAccountTransaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinaccounttransaction.FieldAmount, coinaccounttransaction.FieldCreateAt, coinaccounttransaction.FieldUpdateAt, coinaccounttransaction.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case coinaccounttransaction.FieldMessage, coinaccounttransaction.FieldState, coinaccounttransaction.FieldChainTransactionID:
			values[i] = new(sql.NullString)
		case coinaccounttransaction.FieldID, coinaccounttransaction.FieldUserID, coinaccounttransaction.FieldAppID, coinaccounttransaction.FieldFromAddressID, coinaccounttransaction.FieldToAddressID, coinaccounttransaction.FieldCoinTypeID, coinaccounttransaction.FieldPlatformTransactionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinAccountTransaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinAccountTransaction fields.
func (cat *CoinAccountTransaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinaccounttransaction.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cat.ID = *value
			}
		case coinaccounttransaction.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				cat.UserID = *value
			}
		case coinaccounttransaction.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				cat.AppID = *value
			}
		case coinaccounttransaction.FieldFromAddressID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field from_address_id", values[i])
			} else if value != nil {
				cat.FromAddressID = *value
			}
		case coinaccounttransaction.FieldToAddressID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field to_address_id", values[i])
			} else if value != nil {
				cat.ToAddressID = *value
			}
		case coinaccounttransaction.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cat.CoinTypeID = *value
			}
		case coinaccounttransaction.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				cat.Amount = uint64(value.Int64)
			}
		case coinaccounttransaction.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				cat.Message = value.String
			}
		case coinaccounttransaction.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				cat.State = coinaccounttransaction.State(value.String)
			}
		case coinaccounttransaction.FieldChainTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field chain_transaction_id", values[i])
			} else if value.Valid {
				cat.ChainTransactionID = value.String
			}
		case coinaccounttransaction.FieldPlatformTransactionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field platform_transaction_id", values[i])
			} else if value != nil {
				cat.PlatformTransactionID = *value
			}
		case coinaccounttransaction.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				cat.CreateAt = uint32(value.Int64)
			}
		case coinaccounttransaction.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				cat.UpdateAt = uint32(value.Int64)
			}
		case coinaccounttransaction.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				cat.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinAccountTransaction.
// Note that you need to call CoinAccountTransaction.Unwrap() before calling this method if this CoinAccountTransaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (cat *CoinAccountTransaction) Update() *CoinAccountTransactionUpdateOne {
	return (&CoinAccountTransactionClient{config: cat.config}).UpdateOne(cat)
}

// Unwrap unwraps the CoinAccountTransaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cat *CoinAccountTransaction) Unwrap() *CoinAccountTransaction {
	tx, ok := cat.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinAccountTransaction is not a transactional entity")
	}
	cat.config.driver = tx.drv
	return cat
}

// String implements the fmt.Stringer.
func (cat *CoinAccountTransaction) String() string {
	var builder strings.Builder
	builder.WriteString("CoinAccountTransaction(")
	builder.WriteString(fmt.Sprintf("id=%v", cat.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", cat.UserID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", cat.AppID))
	builder.WriteString(", from_address_id=")
	builder.WriteString(fmt.Sprintf("%v", cat.FromAddressID))
	builder.WriteString(", to_address_id=")
	builder.WriteString(fmt.Sprintf("%v", cat.ToAddressID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cat.CoinTypeID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", cat.Amount))
	builder.WriteString(", message=")
	builder.WriteString(cat.Message)
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", cat.State))
	builder.WriteString(", chain_transaction_id=")
	builder.WriteString(cat.ChainTransactionID)
	builder.WriteString(", platform_transaction_id=")
	builder.WriteString(fmt.Sprintf("%v", cat.PlatformTransactionID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", cat.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", cat.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", cat.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// CoinAccountTransactions is a parsable slice of CoinAccountTransaction.
type CoinAccountTransactions []*CoinAccountTransaction

func (cat CoinAccountTransactions) config(cfg config) {
	for _i := range cat {
		cat[_i].config = cfg
	}
}