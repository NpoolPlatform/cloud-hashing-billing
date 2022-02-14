// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/google/uuid"
)

// PlatformSetting is the model entity for the PlatformSetting schema.
type PlatformSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// WarmAccountUsdAmount holds the value of the "warm_account_usd_amount" field.
	WarmAccountUsdAmount uint64 `json:"warm_account_usd_amount,omitempty"`
	// PaymentAccountUsdAmount holds the value of the "payment_account_usd_amount" field.
	PaymentAccountUsdAmount uint64 `json:"payment_account_usd_amount,omitempty"`
	// WithdrawAutoReviewUsdAmount holds the value of the "withdraw_auto_review_usd_amount" field.
	WithdrawAutoReviewUsdAmount uint64 `json:"withdraw_auto_review_usd_amount,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlatformSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case platformsetting.FieldWarmAccountUsdAmount, platformsetting.FieldPaymentAccountUsdAmount, platformsetting.FieldWithdrawAutoReviewUsdAmount, platformsetting.FieldCreateAt, platformsetting.FieldUpdateAt, platformsetting.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case platformsetting.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PlatformSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlatformSetting fields.
func (ps *PlatformSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case platformsetting.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ps.ID = *value
			}
		case platformsetting.FieldWarmAccountUsdAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field warm_account_usd_amount", values[i])
			} else if value.Valid {
				ps.WarmAccountUsdAmount = uint64(value.Int64)
			}
		case platformsetting.FieldPaymentAccountUsdAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_account_usd_amount", values[i])
			} else if value.Valid {
				ps.PaymentAccountUsdAmount = uint64(value.Int64)
			}
		case platformsetting.FieldWithdrawAutoReviewUsdAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_auto_review_usd_amount", values[i])
			} else if value.Valid {
				ps.WithdrawAutoReviewUsdAmount = uint64(value.Int64)
			}
		case platformsetting.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ps.CreateAt = uint32(value.Int64)
			}
		case platformsetting.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				ps.UpdateAt = uint32(value.Int64)
			}
		case platformsetting.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				ps.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PlatformSetting.
// Note that you need to call PlatformSetting.Unwrap() before calling this method if this PlatformSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (ps *PlatformSetting) Update() *PlatformSettingUpdateOne {
	return (&PlatformSettingClient{config: ps.config}).UpdateOne(ps)
}

// Unwrap unwraps the PlatformSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ps *PlatformSetting) Unwrap() *PlatformSetting {
	tx, ok := ps.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlatformSetting is not a transactional entity")
	}
	ps.config.driver = tx.drv
	return ps
}

// String implements the fmt.Stringer.
func (ps *PlatformSetting) String() string {
	var builder strings.Builder
	builder.WriteString("PlatformSetting(")
	builder.WriteString(fmt.Sprintf("id=%v", ps.ID))
	builder.WriteString(", warm_account_usd_amount=")
	builder.WriteString(fmt.Sprintf("%v", ps.WarmAccountUsdAmount))
	builder.WriteString(", payment_account_usd_amount=")
	builder.WriteString(fmt.Sprintf("%v", ps.PaymentAccountUsdAmount))
	builder.WriteString(", withdraw_auto_review_usd_amount=")
	builder.WriteString(fmt.Sprintf("%v", ps.WithdrawAutoReviewUsdAmount))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ps.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", ps.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", ps.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// PlatformSettings is a parsable slice of PlatformSetting.
type PlatformSettings []*PlatformSetting

func (ps PlatformSettings) config(cfg config) {
	for _i := range ps {
		ps[_i].config = cfg
	}
}
