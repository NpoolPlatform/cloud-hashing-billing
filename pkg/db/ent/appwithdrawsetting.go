// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/appwithdrawsetting"
	"github.com/google/uuid"
)

// AppWithdrawSetting is the model entity for the AppWithdrawSetting schema.
type AppWithdrawSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// WithdrawAutoReviewCoinAmount holds the value of the "withdraw_auto_review_coin_amount" field.
	WithdrawAutoReviewCoinAmount uint64 `json:"withdraw_auto_review_coin_amount,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppWithdrawSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appwithdrawsetting.FieldWithdrawAutoReviewCoinAmount, appwithdrawsetting.FieldCreateAt, appwithdrawsetting.FieldUpdateAt, appwithdrawsetting.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case appwithdrawsetting.FieldID, appwithdrawsetting.FieldAppID, appwithdrawsetting.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppWithdrawSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppWithdrawSetting fields.
func (aws *AppWithdrawSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appwithdrawsetting.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				aws.ID = *value
			}
		case appwithdrawsetting.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				aws.AppID = *value
			}
		case appwithdrawsetting.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				aws.CoinTypeID = *value
			}
		case appwithdrawsetting.FieldWithdrawAutoReviewCoinAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field withdraw_auto_review_coin_amount", values[i])
			} else if value.Valid {
				aws.WithdrawAutoReviewCoinAmount = uint64(value.Int64)
			}
		case appwithdrawsetting.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				aws.CreateAt = uint32(value.Int64)
			}
		case appwithdrawsetting.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				aws.UpdateAt = uint32(value.Int64)
			}
		case appwithdrawsetting.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				aws.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppWithdrawSetting.
// Note that you need to call AppWithdrawSetting.Unwrap() before calling this method if this AppWithdrawSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (aws *AppWithdrawSetting) Update() *AppWithdrawSettingUpdateOne {
	return (&AppWithdrawSettingClient{config: aws.config}).UpdateOne(aws)
}

// Unwrap unwraps the AppWithdrawSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aws *AppWithdrawSetting) Unwrap() *AppWithdrawSetting {
	_tx, ok := aws.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppWithdrawSetting is not a transactional entity")
	}
	aws.config.driver = _tx.drv
	return aws
}

// String implements the fmt.Stringer.
func (aws *AppWithdrawSetting) String() string {
	var builder strings.Builder
	builder.WriteString("AppWithdrawSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aws.ID))
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", aws.AppID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", aws.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("withdraw_auto_review_coin_amount=")
	builder.WriteString(fmt.Sprintf("%v", aws.WithdrawAutoReviewCoinAmount))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", aws.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", aws.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(fmt.Sprintf("%v", aws.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// AppWithdrawSettings is a parsable slice of AppWithdrawSetting.
type AppWithdrawSettings []*AppWithdrawSetting

func (aws AppWithdrawSettings) config(cfg config) {
	for _i := range aws {
		aws[_i].config = cfg
	}
}
