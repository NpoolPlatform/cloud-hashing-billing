// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/appwithdrawsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccountinfo"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinaccounttransaction"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/coinsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/goodpayment"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformsetting"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/schema"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userdirectbenefit"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdraw"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/userwithdrawitem"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appwithdrawsettingFields := schema.AppWithdrawSetting{}.Fields()
	_ = appwithdrawsettingFields
	// appwithdrawsettingDescCreateAt is the schema descriptor for create_at field.
	appwithdrawsettingDescCreateAt := appwithdrawsettingFields[4].Descriptor()
	// appwithdrawsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	appwithdrawsetting.DefaultCreateAt = appwithdrawsettingDescCreateAt.Default.(func() uint32)
	// appwithdrawsettingDescUpdateAt is the schema descriptor for update_at field.
	appwithdrawsettingDescUpdateAt := appwithdrawsettingFields[5].Descriptor()
	// appwithdrawsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	appwithdrawsetting.DefaultUpdateAt = appwithdrawsettingDescUpdateAt.Default.(func() uint32)
	// appwithdrawsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appwithdrawsetting.UpdateDefaultUpdateAt = appwithdrawsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// appwithdrawsettingDescDeleteAt is the schema descriptor for delete_at field.
	appwithdrawsettingDescDeleteAt := appwithdrawsettingFields[6].Descriptor()
	// appwithdrawsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	appwithdrawsetting.DefaultDeleteAt = appwithdrawsettingDescDeleteAt.Default.(func() uint32)
	// appwithdrawsettingDescID is the schema descriptor for id field.
	appwithdrawsettingDescID := appwithdrawsettingFields[0].Descriptor()
	// appwithdrawsetting.DefaultID holds the default value on creation for the id field.
	appwithdrawsetting.DefaultID = appwithdrawsettingDescID.Default.(func() uuid.UUID)
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
	coinsettingDescCreateAt := coinsettingFields[8].Descriptor()
	// coinsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	coinsetting.DefaultCreateAt = coinsettingDescCreateAt.Default.(func() uint32)
	// coinsettingDescUpdateAt is the schema descriptor for update_at field.
	coinsettingDescUpdateAt := coinsettingFields[9].Descriptor()
	// coinsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	coinsetting.DefaultUpdateAt = coinsettingDescUpdateAt.Default.(func() uint32)
	// coinsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	coinsetting.UpdateDefaultUpdateAt = coinsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// coinsettingDescDeleteAt is the schema descriptor for delete_at field.
	coinsettingDescDeleteAt := coinsettingFields[10].Descriptor()
	// coinsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	coinsetting.DefaultDeleteAt = coinsettingDescDeleteAt.Default.(func() uint32)
	// coinsettingDescID is the schema descriptor for id field.
	coinsettingDescID := coinsettingFields[0].Descriptor()
	// coinsetting.DefaultID holds the default value on creation for the id field.
	coinsetting.DefaultID = coinsettingDescID.Default.(func() uuid.UUID)
	goodbenefitFields := schema.GoodBenefit{}.Fields()
	_ = goodbenefitFields
	// goodbenefitDescCreateAt is the schema descriptor for create_at field.
	goodbenefitDescCreateAt := goodbenefitFields[4].Descriptor()
	// goodbenefit.DefaultCreateAt holds the default value on creation for the create_at field.
	goodbenefit.DefaultCreateAt = goodbenefitDescCreateAt.Default.(func() uint32)
	// goodbenefitDescUpdateAt is the schema descriptor for update_at field.
	goodbenefitDescUpdateAt := goodbenefitFields[5].Descriptor()
	// goodbenefit.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodbenefit.DefaultUpdateAt = goodbenefitDescUpdateAt.Default.(func() uint32)
	// goodbenefit.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodbenefit.UpdateDefaultUpdateAt = goodbenefitDescUpdateAt.UpdateDefault.(func() uint32)
	// goodbenefitDescDeleteAt is the schema descriptor for delete_at field.
	goodbenefitDescDeleteAt := goodbenefitFields[6].Descriptor()
	// goodbenefit.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodbenefit.DefaultDeleteAt = goodbenefitDescDeleteAt.Default.(func() uint32)
	// goodbenefitDescID is the schema descriptor for id field.
	goodbenefitDescID := goodbenefitFields[0].Descriptor()
	// goodbenefit.DefaultID holds the default value on creation for the id field.
	goodbenefit.DefaultID = goodbenefitDescID.Default.(func() uuid.UUID)
	goodpaymentFields := schema.GoodPayment{}.Fields()
	_ = goodpaymentFields
	// goodpaymentDescCreateAt is the schema descriptor for create_at field.
	goodpaymentDescCreateAt := goodpaymentFields[6].Descriptor()
	// goodpayment.DefaultCreateAt holds the default value on creation for the create_at field.
	goodpayment.DefaultCreateAt = goodpaymentDescCreateAt.Default.(func() uint32)
	// goodpaymentDescUpdateAt is the schema descriptor for update_at field.
	goodpaymentDescUpdateAt := goodpaymentFields[7].Descriptor()
	// goodpayment.DefaultUpdateAt holds the default value on creation for the update_at field.
	goodpayment.DefaultUpdateAt = goodpaymentDescUpdateAt.Default.(func() uint32)
	// goodpayment.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	goodpayment.UpdateDefaultUpdateAt = goodpaymentDescUpdateAt.UpdateDefault.(func() uint32)
	// goodpaymentDescDeleteAt is the schema descriptor for delete_at field.
	goodpaymentDescDeleteAt := goodpaymentFields[8].Descriptor()
	// goodpayment.DefaultDeleteAt holds the default value on creation for the delete_at field.
	goodpayment.DefaultDeleteAt = goodpaymentDescDeleteAt.Default.(func() uint32)
	// goodpaymentDescID is the schema descriptor for id field.
	goodpaymentDescID := goodpaymentFields[0].Descriptor()
	// goodpayment.DefaultID holds the default value on creation for the id field.
	goodpayment.DefaultID = goodpaymentDescID.Default.(func() uuid.UUID)
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
	platformsettingDescCreateAt := platformsettingFields[4].Descriptor()
	// platformsetting.DefaultCreateAt holds the default value on creation for the create_at field.
	platformsetting.DefaultCreateAt = platformsettingDescCreateAt.Default.(func() uint32)
	// platformsettingDescUpdateAt is the schema descriptor for update_at field.
	platformsettingDescUpdateAt := platformsettingFields[5].Descriptor()
	// platformsetting.DefaultUpdateAt holds the default value on creation for the update_at field.
	platformsetting.DefaultUpdateAt = platformsettingDescUpdateAt.Default.(func() uint32)
	// platformsetting.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	platformsetting.UpdateDefaultUpdateAt = platformsettingDescUpdateAt.UpdateDefault.(func() uint32)
	// platformsettingDescDeleteAt is the schema descriptor for delete_at field.
	platformsettingDescDeleteAt := platformsettingFields[6].Descriptor()
	// platformsetting.DefaultDeleteAt holds the default value on creation for the delete_at field.
	platformsetting.DefaultDeleteAt = platformsettingDescDeleteAt.Default.(func() uint32)
	// platformsettingDescID is the schema descriptor for id field.
	platformsettingDescID := platformsettingFields[0].Descriptor()
	// platformsetting.DefaultID holds the default value on creation for the id field.
	platformsetting.DefaultID = platformsettingDescID.Default.(func() uuid.UUID)
	userbenefitFields := schema.UserBenefit{}.Fields()
	_ = userbenefitFields
	// userbenefitDescCreateAt is the schema descriptor for create_at field.
	userbenefitDescCreateAt := userbenefitFields[8].Descriptor()
	// userbenefit.DefaultCreateAt holds the default value on creation for the create_at field.
	userbenefit.DefaultCreateAt = userbenefitDescCreateAt.Default.(func() uint32)
	// userbenefitDescUpdateAt is the schema descriptor for update_at field.
	userbenefitDescUpdateAt := userbenefitFields[9].Descriptor()
	// userbenefit.DefaultUpdateAt holds the default value on creation for the update_at field.
	userbenefit.DefaultUpdateAt = userbenefitDescUpdateAt.Default.(func() uint32)
	// userbenefit.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userbenefit.UpdateDefaultUpdateAt = userbenefitDescUpdateAt.UpdateDefault.(func() uint32)
	// userbenefitDescDeleteAt is the schema descriptor for delete_at field.
	userbenefitDescDeleteAt := userbenefitFields[10].Descriptor()
	// userbenefit.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userbenefit.DefaultDeleteAt = userbenefitDescDeleteAt.Default.(func() uint32)
	// userbenefitDescID is the schema descriptor for id field.
	userbenefitDescID := userbenefitFields[0].Descriptor()
	// userbenefit.DefaultID holds the default value on creation for the id field.
	userbenefit.DefaultID = userbenefitDescID.Default.(func() uuid.UUID)
	userdirectbenefitFields := schema.UserDirectBenefit{}.Fields()
	_ = userdirectbenefitFields
	// userdirectbenefitDescCreateAt is the schema descriptor for create_at field.
	userdirectbenefitDescCreateAt := userdirectbenefitFields[5].Descriptor()
	// userdirectbenefit.DefaultCreateAt holds the default value on creation for the create_at field.
	userdirectbenefit.DefaultCreateAt = userdirectbenefitDescCreateAt.Default.(func() uint32)
	// userdirectbenefitDescUpdateAt is the schema descriptor for update_at field.
	userdirectbenefitDescUpdateAt := userdirectbenefitFields[6].Descriptor()
	// userdirectbenefit.DefaultUpdateAt holds the default value on creation for the update_at field.
	userdirectbenefit.DefaultUpdateAt = userdirectbenefitDescUpdateAt.Default.(func() uint32)
	// userdirectbenefit.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userdirectbenefit.UpdateDefaultUpdateAt = userdirectbenefitDescUpdateAt.UpdateDefault.(func() uint32)
	// userdirectbenefitDescDeleteAt is the schema descriptor for delete_at field.
	userdirectbenefitDescDeleteAt := userdirectbenefitFields[7].Descriptor()
	// userdirectbenefit.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userdirectbenefit.DefaultDeleteAt = userdirectbenefitDescDeleteAt.Default.(func() uint32)
	// userdirectbenefitDescID is the schema descriptor for id field.
	userdirectbenefitDescID := userdirectbenefitFields[0].Descriptor()
	// userdirectbenefit.DefaultID holds the default value on creation for the id field.
	userdirectbenefit.DefaultID = userdirectbenefitDescID.Default.(func() uuid.UUID)
	userwithdrawFields := schema.UserWithdraw{}.Fields()
	_ = userwithdrawFields
	// userwithdrawDescCreateAt is the schema descriptor for create_at field.
	userwithdrawDescCreateAt := userwithdrawFields[7].Descriptor()
	// userwithdraw.DefaultCreateAt holds the default value on creation for the create_at field.
	userwithdraw.DefaultCreateAt = userwithdrawDescCreateAt.Default.(func() uint32)
	// userwithdrawDescUpdateAt is the schema descriptor for update_at field.
	userwithdrawDescUpdateAt := userwithdrawFields[8].Descriptor()
	// userwithdraw.DefaultUpdateAt holds the default value on creation for the update_at field.
	userwithdraw.DefaultUpdateAt = userwithdrawDescUpdateAt.Default.(func() uint32)
	// userwithdraw.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userwithdraw.UpdateDefaultUpdateAt = userwithdrawDescUpdateAt.UpdateDefault.(func() uint32)
	// userwithdrawDescDeleteAt is the schema descriptor for delete_at field.
	userwithdrawDescDeleteAt := userwithdrawFields[9].Descriptor()
	// userwithdraw.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userwithdraw.DefaultDeleteAt = userwithdrawDescDeleteAt.Default.(func() uint32)
	// userwithdrawDescID is the schema descriptor for id field.
	userwithdrawDescID := userwithdrawFields[0].Descriptor()
	// userwithdraw.DefaultID holds the default value on creation for the id field.
	userwithdraw.DefaultID = userwithdrawDescID.Default.(func() uuid.UUID)
	userwithdrawitemFields := schema.UserWithdrawItem{}.Fields()
	_ = userwithdrawitemFields
	// userwithdrawitemDescCreateAt is the schema descriptor for create_at field.
	userwithdrawitemDescCreateAt := userwithdrawitemFields[7].Descriptor()
	// userwithdrawitem.DefaultCreateAt holds the default value on creation for the create_at field.
	userwithdrawitem.DefaultCreateAt = userwithdrawitemDescCreateAt.Default.(func() uint32)
	// userwithdrawitemDescUpdateAt is the schema descriptor for update_at field.
	userwithdrawitemDescUpdateAt := userwithdrawitemFields[8].Descriptor()
	// userwithdrawitem.DefaultUpdateAt holds the default value on creation for the update_at field.
	userwithdrawitem.DefaultUpdateAt = userwithdrawitemDescUpdateAt.Default.(func() uint32)
	// userwithdrawitem.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	userwithdrawitem.UpdateDefaultUpdateAt = userwithdrawitemDescUpdateAt.UpdateDefault.(func() uint32)
	// userwithdrawitemDescDeleteAt is the schema descriptor for delete_at field.
	userwithdrawitemDescDeleteAt := userwithdrawitemFields[9].Descriptor()
	// userwithdrawitem.DefaultDeleteAt holds the default value on creation for the delete_at field.
	userwithdrawitem.DefaultDeleteAt = userwithdrawitemDescDeleteAt.Default.(func() uint32)
	// userwithdrawitemDescID is the schema descriptor for id field.
	userwithdrawitemDescID := userwithdrawitemFields[0].Descriptor()
	// userwithdrawitem.DefaultID holds the default value on creation for the id field.
	userwithdrawitem.DefaultID = userwithdrawitemDescID.Default.(func() uuid.UUID)
}
