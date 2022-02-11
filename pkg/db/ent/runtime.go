// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccounttransaction"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodpayment"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	coinaccountinfoFields := schema.CoinAccountInfo{}.Fields()
	_ = coinaccountinfoFields
	// coinaccountinfoDescCreateAt is the schema descriptor for create_at field.
	coinaccountinfoDescCreateAt := coinaccountinfoFields[4].Descriptor()
	// coinaccountinfo.DefaultCreateAt holds the default value on creation for the create_at field.
	coinaccountinfo.DefaultCreateAt = coinaccountinfoDescCreateAt.Default.(func() uint32)
	// coinaccountinfoDescUpdateAt is the schema descriptor for update_at field.
	coinaccountinfoDescUpdateAt := coinaccountinfoFields[5].Descriptor()
	// coinaccountinfo.DefaultUpdateAt holds the default value on creation for the update_at field.
	coinaccountinfo.DefaultUpdateAt = coinaccountinfoDescUpdateAt.Default.(func() uint32)
	// coinaccountinfo.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	coinaccountinfo.UpdateDefaultUpdateAt = coinaccountinfoDescUpdateAt.UpdateDefault.(func() uint32)
	// coinaccountinfoDescDeleteAt is the schema descriptor for delete_at field.
	coinaccountinfoDescDeleteAt := coinaccountinfoFields[6].Descriptor()
	// coinaccountinfo.DefaultDeleteAt holds the default value on creation for the delete_at field.
	coinaccountinfo.DefaultDeleteAt = coinaccountinfoDescDeleteAt.Default.(func() uint32)
	// coinaccountinfoDescID is the schema descriptor for id field.
	coinaccountinfoDescID := coinaccountinfoFields[0].Descriptor()
	// coinaccountinfo.DefaultID holds the default value on creation for the id field.
	coinaccountinfo.DefaultID = coinaccountinfoDescID.Default.(func() uuid.UUID)
	coinaccounttransactionFields := schema.CoinAccountTransaction{}.Fields()
	_ = coinaccounttransactionFields
	// coinaccounttransactionDescCreateAt is the schema descriptor for create_at field.
	coinaccounttransactionDescCreateAt := coinaccounttransactionFields[10].Descriptor()
	// coinaccounttransaction.DefaultCreateAt holds the default value on creation for the create_at field.
	coinaccounttransaction.DefaultCreateAt = coinaccounttransactionDescCreateAt.Default.(func() uint32)
	// coinaccounttransactionDescUpdateAt is the schema descriptor for update_at field.
	coinaccounttransactionDescUpdateAt := coinaccounttransactionFields[11].Descriptor()
	// coinaccounttransaction.DefaultUpdateAt holds the default value on creation for the update_at field.
	coinaccounttransaction.DefaultUpdateAt = coinaccounttransactionDescUpdateAt.Default.(func() uint32)
	// coinaccounttransaction.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	coinaccounttransaction.UpdateDefaultUpdateAt = coinaccounttransactionDescUpdateAt.UpdateDefault.(func() uint32)
	// coinaccounttransactionDescDeleteAt is the schema descriptor for delete_at field.
	coinaccounttransactionDescDeleteAt := coinaccounttransactionFields[12].Descriptor()
	// coinaccounttransaction.DefaultDeleteAt holds the default value on creation for the delete_at field.
	coinaccounttransaction.DefaultDeleteAt = coinaccounttransactionDescDeleteAt.Default.(func() uint32)
	// coinaccounttransactionDescID is the schema descriptor for id field.
	coinaccounttransactionDescID := coinaccounttransactionFields[0].Descriptor()
	// coinaccounttransaction.DefaultID holds the default value on creation for the id field.
	coinaccounttransaction.DefaultID = coinaccounttransactionDescID.Default.(func() uuid.UUID)
	coinsettingFields := schema.CoinSetting{}.Fields()
	_ = coinsettingFields
	// coinsettingDescCreateAt is the schema descriptor for create_at field.
	coinsettingDescCreateAt := coinsettingFields[3].Descriptor()
	// coinsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	coinsetting.DefaultCreateAt = coinsettingDescCreateAt.Default.(func() uint32)
	// coinsettingDescUpdateAt is the schema descriptor for update_at field.
	coinsettingDescUpdateAt := coinsettingFields[4].Descriptor()
	// coinsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	coinsetting.DefaultUpdateAt = coinsettingDescUpdateAt.Default.(func() uint32)
	// coinsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	coinsetting.UpdateDefaultUpdateAt = coinsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// coinsettingDescDeleteAt is the schema descriptor for delete_at field.
	coinsettingDescDeleteAt := coinsettingFields[5].Descriptor()
	// coinsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	coinsetting.DefaultDeleteAt = coinsettingDescDeleteAt.Default.(func() uint32)
	// coinsettingDescID is the schema descriptor for id field.
	coinsettingDescID := coinsettingFields[0].Descriptor()
	// coinsetting.DefaultID holds the default value on creation for the id field.
	coinsetting.DefaultID = coinsettingDescID.Default.(func() uuid.UUID)
	goodbenefitFields := schema.GoodBenefit{}.Fields()
	_ = goodbenefitFields
	// goodbenefitDescCreateAt is the schema descriptor for create_at field.
	goodbenefitDescCreateAt := goodbenefitFields[7].Descriptor()
	// goodbenefit.DefaultCreateAt holds the default value on creation for the create_at field.
	goodbenefit.DefaultCreateAt = goodbenefitDescCreateAt.Default.(func() uint32)
	// goodbenefitDescUpdateAt is the schema descriptor for update_at field.
	goodbenefitDescUpdateAt := goodbenefitFields[8].Descriptor()
	// goodbenefit.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodbenefit.DefaultUpdateAt = goodbenefitDescUpdateAt.Default.(func() uint32)
	// goodbenefit.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodbenefit.UpdateDefaultUpdateAt = goodbenefitDescUpdateAt.UpdateDefault.(func() uint32)
	// goodbenefitDescDeleteAt is the schema descriptor for delete_at field.
	goodbenefitDescDeleteAt := goodbenefitFields[9].Descriptor()
	// goodbenefit.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodbenefit.DefaultDeleteAt = goodbenefitDescDeleteAt.Default.(func() uint32)
	// goodbenefitDescID is the schema descriptor for id field.
	goodbenefitDescID := goodbenefitFields[0].Descriptor()
	// goodbenefit.DefaultID holds the default value on creation for the id field.
	goodbenefit.DefaultID = goodbenefitDescID.Default.(func() uuid.UUID)
	goodpaymentFields := schema.GoodPayment{}.Fields()
	_ = goodpaymentFields
	// goodpaymentDescCreateAt is the schema descriptor for create_at field.
	goodpaymentDescCreateAt := goodpaymentFields[5].Descriptor()
	// goodpayment.DefaultCreateAt holds the default value on creation for the create_at field.
	goodpayment.DefaultCreateAt = goodpaymentDescCreateAt.Default.(func() uint32)
	// goodpaymentDescUpdateAt is the schema descriptor for update_at field.
	goodpaymentDescUpdateAt := goodpaymentFields[6].Descriptor()
	// goodpayment.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodpayment.DefaultUpdateAt = goodpaymentDescUpdateAt.Default.(func() uint32)
	// goodpayment.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodpayment.UpdateDefaultUpdateAt = goodpaymentDescUpdateAt.UpdateDefault.(func() uint32)
	// goodpaymentDescDeleteAt is the schema descriptor for delete_at field.
	goodpaymentDescDeleteAt := goodpaymentFields[7].Descriptor()
	// goodpayment.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodpayment.DefaultDeleteAt = goodpaymentDescDeleteAt.Default.(func() uint32)
	// goodpaymentDescID is the schema descriptor for id field.
	goodpaymentDescID := goodpaymentFields[0].Descriptor()
	// goodpayment.DefaultID holds the default value on creation for the id field.
	goodpayment.DefaultID = goodpaymentDescID.Default.(func() uuid.UUID)
	goodsettingFields := schema.GoodSetting{}.Fields()
	_ = goodsettingFields
	// goodsettingDescCreateAt is the schema descriptor for create_at field.
	goodsettingDescCreateAt := goodsettingFields[4].Descriptor()
	// goodsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	goodsetting.DefaultCreateAt = goodsettingDescCreateAt.Default.(func() uint32)
	// goodsettingDescUpdateAt is the schema descriptor for update_at field.
	goodsettingDescUpdateAt := goodsettingFields[5].Descriptor()
	// goodsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodsetting.DefaultUpdateAt = goodsettingDescUpdateAt.Default.(func() uint32)
	// goodsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodsetting.UpdateDefaultUpdateAt = goodsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// goodsettingDescDeleteAt is the schema descriptor for delete_at field.
	goodsettingDescDeleteAt := goodsettingFields[6].Descriptor()
	// goodsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodsetting.DefaultDeleteAt = goodsettingDescDeleteAt.Default.(func() uint32)
	// goodsettingDescID is the schema descriptor for id field.
	goodsettingDescID := goodsettingFields[0].Descriptor()
	// goodsetting.DefaultID holds the default value on creation for the id field.
	goodsetting.DefaultID = goodsettingDescID.Default.(func() uuid.UUID)
	platformbenefitFields := schema.PlatformBenefit{}.Fields()
	_ = platformbenefitFields
	// platformbenefitDescCreateAt is the schema descriptor for create_at field.
	platformbenefitDescCreateAt := platformbenefitFields[6].Descriptor()
	// platformbenefit.DefaultCreateAt holds the default value on creation for the create_at field.
	platformbenefit.DefaultCreateAt = platformbenefitDescCreateAt.Default.(func() uint32)
	// platformbenefitDescUpdateAt is the schema descriptor for update_at field.
	platformbenefitDescUpdateAt := platformbenefitFields[7].Descriptor()
	// platformbenefit.DefaultUpdateAt holds the default value on creation for the update_at field.
	platformbenefit.DefaultUpdateAt = platformbenefitDescUpdateAt.Default.(func() uint32)
	// platformbenefit.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	platformbenefit.UpdateDefaultUpdateAt = platformbenefitDescUpdateAt.UpdateDefault.(func() uint32)
	// platformbenefitDescDeleteAt is the schema descriptor for delete_at field.
	platformbenefitDescDeleteAt := platformbenefitFields[8].Descriptor()
	// platformbenefit.DefaultDeleteAt holds the default value on creation for the delete_at field.
	platformbenefit.DefaultDeleteAt = platformbenefitDescDeleteAt.Default.(func() uint32)
	// platformbenefitDescID is the schema descriptor for id field.
	platformbenefitDescID := platformbenefitFields[0].Descriptor()
	// platformbenefit.DefaultID holds the default value on creation for the id field.
	platformbenefit.DefaultID = platformbenefitDescID.Default.(func() uuid.UUID)
	platformsettingFields := schema.PlatformSetting{}.Fields()
	_ = platformsettingFields
	// platformsettingDescCreateAt is the schema descriptor for create_at field.
	platformsettingDescCreateAt := platformsettingFields[2].Descriptor()
	// platformsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	platformsetting.DefaultCreateAt = platformsettingDescCreateAt.Default.(func() uint32)
	// platformsettingDescUpdateAt is the schema descriptor for update_at field.
	platformsettingDescUpdateAt := platformsettingFields[3].Descriptor()
	// platformsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	platformsetting.DefaultUpdateAt = platformsettingDescUpdateAt.Default.(func() uint32)
	// platformsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	platformsetting.UpdateDefaultUpdateAt = platformsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// platformsettingDescDeleteAt is the schema descriptor for delete_at field.
	platformsettingDescDeleteAt := platformsettingFields[4].Descriptor()
	// platformsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	platformsetting.DefaultDeleteAt = platformsettingDescDeleteAt.Default.(func() uint32)
	// platformsettingDescID is the schema descriptor for id field.
	platformsettingDescID := platformsettingFields[0].Descriptor()
	// platformsetting.DefaultID holds the default value on creation for the id field.
	platformsetting.DefaultID = platformsettingDescID.Default.(func() uuid.UUID)
	userbenefitFields := schema.UserBenefit{}.Fields()
	_ = userbenefitFields
	// userbenefitDescCreateAt is the schema descriptor for create_at field.
	userbenefitDescCreateAt := userbenefitFields[7].Descriptor()
	// userbenefit.DefaultCreateAt holds the default value on creation for the create_at field.
	userbenefit.DefaultCreateAt = userbenefitDescCreateAt.Default.(func() uint32)
	// userbenefitDescUpdateAt is the schema descriptor for update_at field.
	userbenefitDescUpdateAt := userbenefitFields[8].Descriptor()
	// userbenefit.DefaultUpdateAt holds the default value on creation for the update_at field.
	userbenefit.DefaultUpdateAt = userbenefitDescUpdateAt.Default.(func() uint32)
	// userbenefit.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userbenefit.UpdateDefaultUpdateAt = userbenefitDescUpdateAt.UpdateDefault.(func() uint32)
	// userbenefitDescDeleteAt is the schema descriptor for delete_at field.
	userbenefitDescDeleteAt := userbenefitFields[9].Descriptor()
	// userbenefit.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userbenefit.DefaultDeleteAt = userbenefitDescDeleteAt.Default.(func() uint32)
	// userbenefitDescID is the schema descriptor for id field.
	userbenefitDescID := userbenefitFields[0].Descriptor()
	// userbenefit.DefaultID holds the default value on creation for the id field.
	userbenefit.DefaultID = userbenefitDescID.Default.(func() uuid.UUID)
	userwithdrawFields := schema.UserWithdraw{}.Fields()
	_ = userwithdrawFields
	// userwithdrawDescCreateAt is the schema descriptor for create_at field.
	userwithdrawDescCreateAt := userwithdrawFields[6].Descriptor()
	// userwithdraw.DefaultCreateAt holds the default value on creation for the create_at field.
	userwithdraw.DefaultCreateAt = userwithdrawDescCreateAt.Default.(func() uint32)
	// userwithdrawDescUpdateAt is the schema descriptor for update_at field.
	userwithdrawDescUpdateAt := userwithdrawFields[7].Descriptor()
	// userwithdraw.DefaultUpdateAt holds the default value on creation for the update_at field.
	userwithdraw.DefaultUpdateAt = userwithdrawDescUpdateAt.Default.(func() uint32)
	// userwithdraw.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userwithdraw.UpdateDefaultUpdateAt = userwithdrawDescUpdateAt.UpdateDefault.(func() uint32)
	// userwithdrawDescDeleteAt is the schema descriptor for delete_at field.
	userwithdrawDescDeleteAt := userwithdrawFields[8].Descriptor()
	// userwithdraw.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userwithdraw.DefaultDeleteAt = userwithdrawDescDeleteAt.Default.(func() uint32)
	// userwithdrawDescID is the schema descriptor for id field.
	userwithdrawDescID := userwithdrawFields[0].Descriptor()
	// userwithdraw.DefaultID holds the default value on creation for the id field.
	userwithdraw.DefaultID = userwithdrawDescID.Default.(func() uuid.UUID)
}
