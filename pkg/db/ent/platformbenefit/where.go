// Code generated by ent, DO NOT EDIT.

package platformbenefit

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// BenefitAccountID applies equality check predicate on the "benefit_account_id" field. It's identical to BenefitAccountIDEQ.
func BenefitAccountID(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitAccountID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// LastBenefitTimestamp applies equality check predicate on the "last_benefit_timestamp" field. It's identical to LastBenefitTimestampEQ.
func LastBenefitTimestamp(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastBenefitTimestamp), v))
	})
}

// ChainTransactionID applies equality check predicate on the "chain_transaction_id" field. It's identical to ChainTransactionIDEQ.
func ChainTransactionID(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainTransactionID), v))
	})
}

// CreateAt applies equality check predicate on the "create_at" field. It's identical to CreateAtEQ.
func CreateAt(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// UpdateAt applies equality check predicate on the "update_at" field. It's identical to UpdateAtEQ.
func UpdateAt(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// BenefitAccountIDEQ applies the EQ predicate on the "benefit_account_id" field.
func BenefitAccountIDEQ(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitAccountID), v))
	})
}

// BenefitAccountIDNEQ applies the NEQ predicate on the "benefit_account_id" field.
func BenefitAccountIDNEQ(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefitAccountID), v))
	})
}

// BenefitAccountIDIn applies the In predicate on the "benefit_account_id" field.
func BenefitAccountIDIn(vs ...uuid.UUID) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBenefitAccountID), v...))
	})
}

// BenefitAccountIDNotIn applies the NotIn predicate on the "benefit_account_id" field.
func BenefitAccountIDNotIn(vs ...uuid.UUID) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBenefitAccountID), v...))
	})
}

// BenefitAccountIDGT applies the GT predicate on the "benefit_account_id" field.
func BenefitAccountIDGT(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefitAccountID), v))
	})
}

// BenefitAccountIDGTE applies the GTE predicate on the "benefit_account_id" field.
func BenefitAccountIDGTE(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefitAccountID), v))
	})
}

// BenefitAccountIDLT applies the LT predicate on the "benefit_account_id" field.
func BenefitAccountIDLT(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefitAccountID), v))
	})
}

// BenefitAccountIDLTE applies the LTE predicate on the "benefit_account_id" field.
func BenefitAccountIDLTE(v uuid.UUID) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefitAccountID), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...uint64) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func AmountNotIn(vs ...uint64) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func AmountGT(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v uint64) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// LastBenefitTimestampEQ applies the EQ predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastBenefitTimestamp), v))
	})
}

// LastBenefitTimestampNEQ applies the NEQ predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampNEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastBenefitTimestamp), v))
	})
}

// LastBenefitTimestampIn applies the In predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastBenefitTimestamp), v...))
	})
}

// LastBenefitTimestampNotIn applies the NotIn predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampNotIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastBenefitTimestamp), v...))
	})
}

// LastBenefitTimestampGT applies the GT predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampGT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastBenefitTimestamp), v))
	})
}

// LastBenefitTimestampGTE applies the GTE predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampGTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastBenefitTimestamp), v))
	})
}

// LastBenefitTimestampLT applies the LT predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampLT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastBenefitTimestamp), v))
	})
}

// LastBenefitTimestampLTE applies the LTE predicate on the "last_benefit_timestamp" field.
func LastBenefitTimestampLTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastBenefitTimestamp), v))
	})
}

// ChainTransactionIDEQ applies the EQ predicate on the "chain_transaction_id" field.
func ChainTransactionIDEQ(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDNEQ applies the NEQ predicate on the "chain_transaction_id" field.
func ChainTransactionIDNEQ(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDIn applies the In predicate on the "chain_transaction_id" field.
func ChainTransactionIDIn(vs ...string) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldChainTransactionID), v...))
	})
}

// ChainTransactionIDNotIn applies the NotIn predicate on the "chain_transaction_id" field.
func ChainTransactionIDNotIn(vs ...string) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldChainTransactionID), v...))
	})
}

// ChainTransactionIDGT applies the GT predicate on the "chain_transaction_id" field.
func ChainTransactionIDGT(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDGTE applies the GTE predicate on the "chain_transaction_id" field.
func ChainTransactionIDGTE(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDLT applies the LT predicate on the "chain_transaction_id" field.
func ChainTransactionIDLT(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDLTE applies the LTE predicate on the "chain_transaction_id" field.
func ChainTransactionIDLTE(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDContains applies the Contains predicate on the "chain_transaction_id" field.
func ChainTransactionIDContains(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDHasPrefix applies the HasPrefix predicate on the "chain_transaction_id" field.
func ChainTransactionIDHasPrefix(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDHasSuffix applies the HasSuffix predicate on the "chain_transaction_id" field.
func ChainTransactionIDHasSuffix(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDEqualFold applies the EqualFold predicate on the "chain_transaction_id" field.
func ChainTransactionIDEqualFold(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChainTransactionID), v))
	})
}

// ChainTransactionIDContainsFold applies the ContainsFold predicate on the "chain_transaction_id" field.
func ChainTransactionIDContainsFold(v string) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChainTransactionID), v))
	})
}

// CreateAtEQ applies the EQ predicate on the "create_at" field.
func CreateAtEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtNEQ applies the NEQ predicate on the "create_at" field.
func CreateAtNEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateAt), v))
	})
}

// CreateAtIn applies the In predicate on the "create_at" field.
func CreateAtIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func CreateAtNotIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func CreateAtGT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateAt), v))
	})
}

// CreateAtGTE applies the GTE predicate on the "create_at" field.
func CreateAtGTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateAt), v))
	})
}

// CreateAtLT applies the LT predicate on the "create_at" field.
func CreateAtLT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateAt), v))
	})
}

// CreateAtLTE applies the LTE predicate on the "create_at" field.
func CreateAtLTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateAt), v))
	})
}

// UpdateAtEQ applies the EQ predicate on the "update_at" field.
func UpdateAtEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtNEQ applies the NEQ predicate on the "update_at" field.
func UpdateAtNEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtIn applies the In predicate on the "update_at" field.
func UpdateAtIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func UpdateAtNotIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func UpdateAtGT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtGTE applies the GTE predicate on the "update_at" field.
func UpdateAtGTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLT applies the LT predicate on the "update_at" field.
func UpdateAtLT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateAt), v))
	})
}

// UpdateAtLTE applies the LTE predicate on the "update_at" field.
func UpdateAtLTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateAt), v))
	})
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func DeleteAtNotIn(vs ...uint32) predicate.PlatformBenefit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func DeleteAtGT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteAt), v))
	})
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v uint32) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlatformBenefit) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlatformBenefit) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
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
func Not(p predicate.PlatformBenefit) predicate.PlatformBenefit {
	return predicate.PlatformBenefit(func(s *sql.Selector) {
		p(s.Not())
	})
}
