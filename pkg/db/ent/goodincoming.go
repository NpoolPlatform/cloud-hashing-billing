// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodincoming"
	"github.com/google/uuid"
)

// GoodIncoming is the model entity for the GoodIncoming schema.
type GoodIncoming struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
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
func (*GoodIncoming) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodincoming.FieldCreateAt, goodincoming.FieldUpdateAt, goodincoming.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case goodincoming.FieldID, goodincoming.FieldGoodID, goodincoming.FieldCoinTypeID, goodincoming.FieldAccountID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodIncoming", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodIncoming fields.
func (gi *GoodIncoming) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodincoming.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gi.ID = *value
			}
		case goodincoming.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gi.GoodID = *value
			}
		case goodincoming.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				gi.CoinTypeID = *value
			}
		case goodincoming.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				gi.AccountID = *value
			}
		case goodincoming.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				gi.CreateAt = uint32(value.Int64)
			}
		case goodincoming.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				gi.UpdateAt = uint32(value.Int64)
			}
		case goodincoming.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				gi.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodIncoming.
// Note that you need to call GoodIncoming.Unwrap() before calling this method if this GoodIncoming
// was returned from a transaction, and the transaction was committed or rolled back.
func (gi *GoodIncoming) Update() *GoodIncomingUpdateOne {
	return (&GoodIncomingClient{config: gi.config}).UpdateOne(gi)
}

// Unwrap unwraps the GoodIncoming entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gi *GoodIncoming) Unwrap() *GoodIncoming {
	tx, ok := gi.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodIncoming is not a transactional entity")
	}
	gi.config.driver = tx.drv
	return gi
}

// String implements the fmt.Stringer.
func (gi *GoodIncoming) String() string {
	var builder strings.Builder
	builder.WriteString("GoodIncoming(")
	builder.WriteString(fmt.Sprintf("id=%v", gi.ID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", gi.GoodID))
	builder.WriteString(", coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", gi.CoinTypeID))
	builder.WriteString(", account_id=")
	builder.WriteString(fmt.Sprintf("%v", gi.AccountID))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", gi.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", gi.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", gi.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// GoodIncomings is a parsable slice of GoodIncoming.
type GoodIncomings []*GoodIncoming

func (gi GoodIncomings) config(cfg config) {
	for _i := range gi {
		gi[_i].config = cfg
	}
}