// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"
	"github.com/google/uuid"
)

// UserWithdraw is the model entity for the UserWithdraw schema.
type UserWithdraw struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels []string `json:"labels,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserWithdraw) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userwithdraw.FieldLabels:
			values[i] = new([]byte)
		case userwithdraw.FieldCreateAt, userwithdraw.FieldUpdateAt, userwithdraw.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case userwithdraw.FieldName, userwithdraw.FieldMessage:
			values[i] = new(sql.NullString)
		case userwithdraw.FieldID, userwithdraw.FieldAppID, userwithdraw.FieldUserID, userwithdraw.FieldCoinTypeID, userwithdraw.FieldAccountID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserWithdraw", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserWithdraw fields.
func (uw *UserWithdraw) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userwithdraw.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				uw.ID = *value
			}
		case userwithdraw.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				uw.AppID = *value
			}
		case userwithdraw.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				uw.UserID = *value
			}
		case userwithdraw.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				uw.Name = value.String
			}
		case userwithdraw.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				uw.Message = value.String
			}
		case userwithdraw.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &uw.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case userwithdraw.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				uw.CoinTypeID = *value
			}
		case userwithdraw.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				uw.AccountID = *value
			}
		case userwithdraw.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				uw.CreateAt = uint32(value.Int64)
			}
		case userwithdraw.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				uw.UpdateAt = uint32(value.Int64)
			}
		case userwithdraw.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				uw.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this UserWithdraw.
// Note that you need to call UserWithdraw.Unwrap() before calling this method if this UserWithdraw
// was returned from a transaction, and the transaction was committed or rolled back.
func (uw *UserWithdraw) Update() *UserWithdrawUpdateOne {
	return (&UserWithdrawClient{config: uw.config}).UpdateOne(uw)
}

// Unwrap unwraps the UserWithdraw entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uw *UserWithdraw) Unwrap() *UserWithdraw {
	_tx, ok := uw.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserWithdraw is not a transactional entity")
	}
	uw.config.driver = _tx.drv
	return uw
}

// String implements the fmt.Stringer.
func (uw *UserWithdraw) String() string {
	var builder strings.Builder
	builder.WriteString("UserWithdraw(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uw.ID))
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", uw.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uw.UserID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(uw.Name)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(uw.Message)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", uw.Labels))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", uw.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", uw.AccountID))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", uw.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", uw.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(fmt.Sprintf("%v", uw.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// UserWithdraws is a parsable slice of UserWithdraw.
type UserWithdraws []*UserWithdraw

func (uw UserWithdraws) config(cfg config) {
	for _i := range uw {
		uw[_i].config = cfg
	}
}
