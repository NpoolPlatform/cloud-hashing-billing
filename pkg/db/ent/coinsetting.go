// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinsetting"
	"github.com/google/uuid"
)

// CoinSetting is the model entity for the CoinSetting schema.
type CoinSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// WarmAccountCoinAmount holds the value of the "warm_account_coin_amount" field.
	WarmAccountCoinAmount uint64 `json:"warm_account_coin_amount,omitempty"`
	// PaymentAccountCoinAmount holds the value of the "payment_account_coin_amount" field.
	PaymentAccountCoinAmount uint64 `json:"payment_account_coin_amount,omitempty"`
	// PlatformOfflineAccountID holds the value of the "platform_offline_account_id" field.
	PlatformOfflineAccountID uuid.UUID `json:"platform_offline_account_id,omitempty"`
	// UserOnlineAccountID holds the value of the "user_online_account_id" field.
	UserOnlineAccountID uuid.UUID `json:"user_online_account_id,omitempty"`
	// UserOfflineAccountID holds the value of the "user_offline_account_id" field.
	UserOfflineAccountID uuid.UUID `json:"user_offline_account_id,omitempty"`
	// GoodIncomingAccountID holds the value of the "good_incoming_account_id" field.
	GoodIncomingAccountID uuid.UUID `json:"good_incoming_account_id,omitempty"`
	// GasProviderAccountID holds the value of the "gas_provider_account_id" field.
	GasProviderAccountID uuid.UUID `json:"gas_provider_account_id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CoinSetting) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case coinsetting.FieldWarmAccountCoinAmount, coinsetting.FieldPaymentAccountCoinAmount, coinsetting.FieldCreateAt, coinsetting.FieldUpdateAt, coinsetting.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case coinsetting.FieldID, coinsetting.FieldCoinTypeID, coinsetting.FieldPlatformOfflineAccountID, coinsetting.FieldUserOnlineAccountID, coinsetting.FieldUserOfflineAccountID, coinsetting.FieldGoodIncomingAccountID, coinsetting.FieldGasProviderAccountID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type CoinSetting", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CoinSetting fields.
func (cs *CoinSetting) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case coinsetting.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cs.ID = *value
			}
		case coinsetting.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				cs.CoinTypeID = *value
			}
		case coinsetting.FieldWarmAccountCoinAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field warm_account_coin_amount", values[i])
			} else if value.Valid {
				cs.WarmAccountCoinAmount = uint64(value.Int64)
			}
		case coinsetting.FieldPaymentAccountCoinAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_account_coin_amount", values[i])
			} else if value.Valid {
				cs.PaymentAccountCoinAmount = uint64(value.Int64)
			}
		case coinsetting.FieldPlatformOfflineAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field platform_offline_account_id", values[i])
			} else if value != nil {
				cs.PlatformOfflineAccountID = *value
			}
		case coinsetting.FieldUserOnlineAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_online_account_id", values[i])
			} else if value != nil {
				cs.UserOnlineAccountID = *value
			}
		case coinsetting.FieldUserOfflineAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_offline_account_id", values[i])
			} else if value != nil {
				cs.UserOfflineAccountID = *value
			}
		case coinsetting.FieldGoodIncomingAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_incoming_account_id", values[i])
			} else if value != nil {
				cs.GoodIncomingAccountID = *value
			}
		case coinsetting.FieldGasProviderAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field gas_provider_account_id", values[i])
			} else if value != nil {
				cs.GasProviderAccountID = *value
			}
		case coinsetting.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				cs.CreateAt = uint32(value.Int64)
			}
		case coinsetting.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				cs.UpdateAt = uint32(value.Int64)
			}
		case coinsetting.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				cs.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this CoinSetting.
// Note that you need to call CoinSetting.Unwrap() before calling this method if this CoinSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *CoinSetting) Update() *CoinSettingUpdateOne {
	return (&CoinSettingClient{config: cs.config}).UpdateOne(cs)
}

// Unwrap unwraps the CoinSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *CoinSetting) Unwrap() *CoinSetting {
	_tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("ent: CoinSetting is not a transactional entity")
	}
	cs.config.driver = _tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *CoinSetting) String() string {
	var builder strings.Builder
	builder.WriteString("CoinSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cs.ID))
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("warm_account_coin_amount=")
	builder.WriteString(fmt.Sprintf("%v", cs.WarmAccountCoinAmount))
	builder.WriteString(", ")
	builder.WriteString("payment_account_coin_amount=")
	builder.WriteString(fmt.Sprintf("%v", cs.PaymentAccountCoinAmount))
	builder.WriteString(", ")
	builder.WriteString("platform_offline_account_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.PlatformOfflineAccountID))
	builder.WriteString(", ")
	builder.WriteString("user_online_account_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.UserOnlineAccountID))
	builder.WriteString(", ")
	builder.WriteString("user_offline_account_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.UserOfflineAccountID))
	builder.WriteString(", ")
	builder.WriteString("good_incoming_account_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.GoodIncomingAccountID))
	builder.WriteString(", ")
	builder.WriteString("gas_provider_account_id=")
	builder.WriteString(fmt.Sprintf("%v", cs.GasProviderAccountID))
	builder.WriteString(", ")
	builder.WriteString("create_at=")
	builder.WriteString(fmt.Sprintf("%v", cs.CreateAt))
	builder.WriteString(", ")
	builder.WriteString("update_at=")
	builder.WriteString(fmt.Sprintf("%v", cs.UpdateAt))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(fmt.Sprintf("%v", cs.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// CoinSettings is a parsable slice of CoinSetting.
type CoinSettings []*CoinSetting

func (cs CoinSettings) config(cfg config) {
	for _i := range cs {
		cs[_i].config = cfg
	}
}
