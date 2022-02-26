// Code generated by entc, DO NOT EDIT.

package userwithdrawitem

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// WithdrawToAccountID applies equality check predicate on the "withdraw_to_account_id" field. It's identical to WithdrawToAccountIDEQ.
func WithdrawToAccountID(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawToAccountID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// PlatformTransactionID applies equality check predicate on the "platform_transaction_id" field. It's identical to PlatformTransactionIDEQ.
func PlatformTransactionID(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformTransactionID), v))
	})
}

// WithdrawType applies equality check predicate on the "withdraw_type" field. It's identical to WithdrawTypeEQ.
func WithdrawType(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawType), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// WithdrawToAccountIDEQ applies the EQ predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawToAccountID), v))
	})
}

// WithdrawToAccountIDNEQ applies the NEQ predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDNEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawToAccountID), v))
	})
}

// WithdrawToAccountIDIn applies the In predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWithdrawToAccountID), v...))
	})
}

// WithdrawToAccountIDNotIn applies the NotIn predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDNotIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWithdrawToAccountID), v...))
	})
}

// WithdrawToAccountIDGT applies the GT predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDGT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawToAccountID), v))
	})
}

// WithdrawToAccountIDGTE applies the GTE predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDGTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawToAccountID), v))
	})
}

// WithdrawToAccountIDLT applies the LT predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDLT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawToAccountID), v))
	})
}

// WithdrawToAccountIDLTE applies the LTE predicate on the "withdraw_to_account_id" field.
func WithdrawToAccountIDLTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawToAccountID), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint64) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...uint64) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint64) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// PlatformTransactionIDEQ applies the EQ predicate on the "platform_transaction_id" field.
func PlatformTransactionIDEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDNEQ applies the NEQ predicate on the "platform_transaction_id" field.
func PlatformTransactionIDNEQ(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDIn applies the In predicate on the "platform_transaction_id" field.
func PlatformTransactionIDIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlatformTransactionID), v...))
	})
}

// PlatformTransactionIDNotIn applies the NotIn predicate on the "platform_transaction_id" field.
func PlatformTransactionIDNotIn(vs ...uuid.UUID) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlatformTransactionID), v...))
	})
}

// PlatformTransactionIDGT applies the GT predicate on the "platform_transaction_id" field.
func PlatformTransactionIDGT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDGTE applies the GTE predicate on the "platform_transaction_id" field.
func PlatformTransactionIDGTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDLT applies the LT predicate on the "platform_transaction_id" field.
func PlatformTransactionIDLT(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlatformTransactionID), v))
	})
}

// PlatformTransactionIDLTE applies the LTE predicate on the "platform_transaction_id" field.
func PlatformTransactionIDLTE(v uuid.UUID) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlatformTransactionID), v))
	})
}

// WithdrawTypeEQ applies the EQ predicate on the "withdraw_type" field.
func WithdrawTypeEQ(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeNEQ applies the NEQ predicate on the "withdraw_type" field.
func WithdrawTypeNEQ(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeIn applies the In predicate on the "withdraw_type" field.
func WithdrawTypeIn(vs ...string) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWithdrawType), v...))
	})
}

// WithdrawTypeNotIn applies the NotIn predicate on the "withdraw_type" field.
func WithdrawTypeNotIn(vs ...string) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWithdrawType), v...))
	})
}

// WithdrawTypeGT applies the GT predicate on the "withdraw_type" field.
func WithdrawTypeGT(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeGTE applies the GTE predicate on the "withdraw_type" field.
func WithdrawTypeGTE(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeLT applies the LT predicate on the "withdraw_type" field.
func WithdrawTypeLT(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeLTE applies the LTE predicate on the "withdraw_type" field.
func WithdrawTypeLTE(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeContains applies the Contains predicate on the "withdraw_type" field.
func WithdrawTypeContains(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeHasPrefix applies the HasPrefix predicate on the "withdraw_type" field.
func WithdrawTypeHasPrefix(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeHasSuffix applies the HasSuffix predicate on the "withdraw_type" field.
func WithdrawTypeHasSuffix(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeEqualFold applies the EqualFold predicate on the "withdraw_type" field.
func WithdrawTypeEqualFold(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWithdrawType), v))
	})
}

// WithdrawTypeContainsFold applies the ContainsFold predicate on the "withdraw_type" field.
func WithdrawTypeContainsFold(v string) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWithdrawType), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateAt), v...))
	})
}

// CreateAtNotIn applies the NotIn predicate on the "create_at" field.
func CreateAtNotIn(vs ...uint32) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateAt), v...))
	})
}

// CreateAtGT applies the GT predicate on the "create_at" field.
func CreateAtGT(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtNotIn applies the NotIn predicate on the "update_at" field.
func UpdateAtNotIn(vs ...uint32) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateAt), v...))
	})
}

// UpdateAtGT applies the GT predicate on the "update_at" field.
func UpdateAtGT(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...uint32) predicate.UserWithdrawItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteAt), v...))
	})
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserWithdrawItem) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserWithdrawItem) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserWithdrawItem) predicate.UserWithdrawItem {
	return predicate.UserWithdrawItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
