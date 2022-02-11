// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodbenefit"
	"github.com/google/uuid"
)

// GoodBenefit is the model entity for the GoodBenefit schema.
type GoodBenefit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// BenefitAccountID holds the value of the "benefit_account_id" field.
	BenefitAccountID uuid.UUID `json:"benefit_account_id,omitempty"`
	// PlatformOfflineAccountID holds the value of the "platform_offline_account_id" field.
	PlatformOfflineAccountID uuid.UUID `json:"platform_offline_account_id,omitempty"`
	// UserOnlineAccountID holds the value of the "user_online_account_id" field.
	UserOnlineAccountID uuid.UUID `json:"user_online_account_id,omitempty"`
	// UserOfflineAccountID holds the value of the "user_offline_account_id" field.
	UserOfflineAccountID uuid.UUID `json:"user_offline_account_id,omitempty"`
	// BenefitIntervalHours holds the value of the "benefit_interval_hours" field.
	BenefitIntervalHours uint32 `json:"benefit_interval_hours,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodBenefit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodbenefit.FieldBenefitIntervalHours, goodbenefit.FieldCreateAt, goodbenefit.FieldUpdateAt, goodbenefit.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case goodbenefit.FieldID, goodbenefit.FieldGoodID, goodbenefit.FieldBenefitAccountID, goodbenefit.FieldPlatformOfflineAccountID, goodbenefit.FieldUserOnlineAccountID, goodbenefit.FieldUserOfflineAccountID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodBenefit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodBenefit fields.
func (gb *GoodBenefit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodbenefit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gb.ID = *value
			}
		case goodbenefit.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gb.GoodID = *value
			}
		case goodbenefit.FieldBenefitAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field benefit_account_id", values[i])
			} else if value != nil {
				gb.BenefitAccountID = *value
			}
		case goodbenefit.FieldPlatformOfflineAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field platform_offline_account_id", values[i])
			} else if value != nil {
				gb.PlatformOfflineAccountID = *value
			}
		case goodbenefit.FieldUserOnlineAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_online_account_id", values[i])
			} else if value != nil {
				gb.UserOnlineAccountID = *value
			}
		case goodbenefit.FieldUserOfflineAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_offline_account_id", values[i])
			} else if value != nil {
				gb.UserOfflineAccountID = *value
			}
		case goodbenefit.FieldBenefitIntervalHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field benefit_interval_hours", values[i])
			} else if value.Valid {
				gb.BenefitIntervalHours = uint32(value.Int64)
			}
		case goodbenefit.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				gb.CreateAt = uint32(value.Int64)
			}
		case goodbenefit.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				gb.UpdateAt = uint32(value.Int64)
			}
		case goodbenefit.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				gb.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodBenefit.
// Note that you need to call GoodBenefit.Unwrap() before calling this method if this GoodBenefit
// was returned from a transaction, and the transaction was committed or rolled back.
func (gb *GoodBenefit) Update() *GoodBenefitUpdateOne {
	return (&GoodBenefitClient{config: gb.config}).UpdateOne(gb)
}

// Unwrap unwraps the GoodBenefit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gb *GoodBenefit) Unwrap() *GoodBenefit {
	tx, ok := gb.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodBenefit is not a transactional entity")
	}
	gb.config.driver = tx.drv
	return gb
}

// String implements the fmt.Stringer.
func (gb *GoodBenefit) String() string {
	var builder strings.Builder
	builder.WriteString("GoodBenefit(")
	builder.WriteString(fmt.Sprintf("id=%v", gb.ID))
	builder.WriteString(", good_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.GoodID))
	builder.WriteString(", benefit_account_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.BenefitAccountID))
	builder.WriteString(", platform_offline_account_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.PlatformOfflineAccountID))
	builder.WriteString(", user_online_account_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.UserOnlineAccountID))
	builder.WriteString(", user_offline_account_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.UserOfflineAccountID))
	builder.WriteString(", benefit_interval_hours=")
	builder.WriteString(fmt.Sprintf("%v", gb.BenefitIntervalHours))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", gb.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", gb.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", gb.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// GoodBenefits is a parsable slice of GoodBenefit.
type GoodBenefits []*GoodBenefit

func (gb GoodBenefits) config(cfg config) {
	for _i := range gb {
		gb[_i].config = cfg
	}
}