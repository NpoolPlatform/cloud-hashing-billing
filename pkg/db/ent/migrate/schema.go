// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinAccountInfosColumns holds the columns for the "coin_account_infos" table.
	CoinAccountInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "address", Type: field.TypeString},
		{Name: "platform_hold_private_key", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CoinAccountInfosTable holds the schema information for the "coin_account_infos" table.
	CoinAccountInfosTable = &schema.Table{
		Name:       "coin_account_infos",
		Columns:    CoinAccountInfosColumns,
		PrimaryKey: []*schema.Column{CoinAccountInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "coinaccountinfo_coin_type_id_address",
				Unique:  true,
				Columns: []*schema.Column{CoinAccountInfosColumns[1], CoinAccountInfosColumns[2]},
			},
		},
	}
	// CoinAccountTransactionsColumns holds the columns for the "coin_account_transactions" table.
	CoinAccountTransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "from_address_id", Type: field.TypeUUID},
		{Name: "to_address_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "message", Type: field.TypeString},
		{Name: "state", Type: field.TypeEnum, Enums: []string{"created", "wait", "paying", "successful", "rejected", "fail"}},
		{Name: "chain_transaction_id", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CoinAccountTransactionsTable holds the schema information for the "coin_account_transactions" table.
	CoinAccountTransactionsTable = &schema.Table{
		Name:       "coin_account_transactions",
		Columns:    CoinAccountTransactionsColumns,
		PrimaryKey: []*schema.Column{CoinAccountTransactionsColumns[0]},
	}
	// CoinSettingsColumns holds the columns for the "coin_settings" table.
	CoinSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Unique: true},
		{Name: "warm_account_coin_amount", Type: field.TypeUint64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// CoinSettingsTable holds the schema information for the "coin_settings" table.
	CoinSettingsTable = &schema.Table{
		Name:       "coin_settings",
		Columns:    CoinSettingsColumns,
		PrimaryKey: []*schema.Column{CoinSettingsColumns[0]},
	}
	// GoodBenefitsColumns holds the columns for the "good_benefits" table.
	GoodBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID, Unique: true},
		{Name: "benefit_account_id", Type: field.TypeUUID},
		{Name: "platform_offline_account_id", Type: field.TypeUUID},
		{Name: "user_online_account_id", Type: field.TypeUUID},
		{Name: "user_offline_account_id", Type: field.TypeUUID},
		{Name: "benefit_interval_hours", Type: field.TypeUint32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// GoodBenefitsTable holds the schema information for the "good_benefits" table.
	GoodBenefitsTable = &schema.Table{
		Name:       "good_benefits",
		Columns:    GoodBenefitsColumns,
		PrimaryKey: []*schema.Column{GoodBenefitsColumns[0]},
	}
	// GoodIncomingsColumns holds the columns for the "good_incomings" table.
	GoodIncomingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// GoodIncomingsTable holds the schema information for the "good_incomings" table.
	GoodIncomingsTable = &schema.Table{
		Name:       "good_incomings",
		Columns:    GoodIncomingsColumns,
		PrimaryKey: []*schema.Column{GoodIncomingsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "goodincoming_good_id_coin_type_id",
				Unique:  true,
				Columns: []*schema.Column{GoodIncomingsColumns[1], GoodIncomingsColumns[2]},
			},
			{
				Name:    "goodincoming_coin_type_id_account_id",
				Unique:  true,
				Columns: []*schema.Column{GoodIncomingsColumns[2], GoodIncomingsColumns[3]},
			},
		},
	}
	// GoodPaymentsColumns holds the columns for the "good_payments" table.
	GoodPaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "payment_coin_type_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID, Unique: true},
		{Name: "idle", Type: field.TypeBool},
		{Name: "occupied_by", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// GoodPaymentsTable holds the schema information for the "good_payments" table.
	GoodPaymentsTable = &schema.Table{
		Name:       "good_payments",
		Columns:    GoodPaymentsColumns,
		PrimaryKey: []*schema.Column{GoodPaymentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "goodpayment_good_id_payment_coin_type_id_account_id",
				Unique:  true,
				Columns: []*schema.Column{GoodPaymentsColumns[1], GoodPaymentsColumns[2], GoodPaymentsColumns[3]},
			},
		},
	}
	// PlatformBenefitsColumns holds the columns for the "platform_benefits" table.
	PlatformBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "benefit_account_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "last_benefit_timestamp", Type: field.TypeUint32, Unique: true},
		{Name: "chain_transaction_id", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PlatformBenefitsTable holds the schema information for the "platform_benefits" table.
	PlatformBenefitsTable = &schema.Table{
		Name:       "platform_benefits",
		Columns:    PlatformBenefitsColumns,
		PrimaryKey: []*schema.Column{PlatformBenefitsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "platformbenefit_good_id_last_benefit_timestamp",
				Unique:  true,
				Columns: []*schema.Column{PlatformBenefitsColumns[1], PlatformBenefitsColumns[4]},
			},
		},
	}
	// PlatformSettingsColumns holds the columns for the "platform_settings" table.
	PlatformSettingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "warm_account_usd_amount", Type: field.TypeUint64},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// PlatformSettingsTable holds the schema information for the "platform_settings" table.
	PlatformSettingsTable = &schema.Table{
		Name:       "platform_settings",
		Columns:    PlatformSettingsColumns,
		PrimaryKey: []*schema.Column{PlatformSettingsColumns[0]},
	}
	// UserBenefitsColumns holds the columns for the "user_benefits" table.
	UserBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "last_benefit_timestamp", Type: field.TypeUint32},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserBenefitsTable holds the schema information for the "user_benefits" table.
	UserBenefitsTable = &schema.Table{
		Name:       "user_benefits",
		Columns:    UserBenefitsColumns,
		PrimaryKey: []*schema.Column{UserBenefitsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userbenefit_good_id_last_benefit_timestamp_app_id_user_id_order_id",
				Unique:  true,
				Columns: []*schema.Column{UserBenefitsColumns[3], UserBenefitsColumns[6], UserBenefitsColumns[1], UserBenefitsColumns[2], UserBenefitsColumns[4]},
			},
		},
	}
	// UserDirectBenefitsColumns holds the columns for the "user_direct_benefits" table.
	UserDirectBenefitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserDirectBenefitsTable holds the schema information for the "user_direct_benefits" table.
	UserDirectBenefitsTable = &schema.Table{
		Name:       "user_direct_benefits",
		Columns:    UserDirectBenefitsColumns,
		PrimaryKey: []*schema.Column{UserDirectBenefitsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userdirectbenefit_app_id_user_id_coin_type_id_account_id",
				Unique:  true,
				Columns: []*schema.Column{UserDirectBenefitsColumns[1], UserDirectBenefitsColumns[2], UserDirectBenefitsColumns[3], UserDirectBenefitsColumns[4]},
			},
		},
	}
	// UserWithdrawsColumns holds the columns for the "user_withdraws" table.
	UserWithdrawsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// UserWithdrawsTable holds the schema information for the "user_withdraws" table.
	UserWithdrawsTable = &schema.Table{
		Name:       "user_withdraws",
		Columns:    UserWithdrawsColumns,
		PrimaryKey: []*schema.Column{UserWithdrawsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "userwithdraw_app_id_user_id_coin_type_id_account_id",
				Unique:  true,
				Columns: []*schema.Column{UserWithdrawsColumns[1], UserWithdrawsColumns[2], UserWithdrawsColumns[5], UserWithdrawsColumns[6]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinAccountInfosTable,
		CoinAccountTransactionsTable,
		CoinSettingsTable,
		GoodBenefitsTable,
		GoodIncomingsTable,
		GoodPaymentsTable,
		PlatformBenefitsTable,
		PlatformSettingsTable,
		UserBenefitsTable,
		UserDirectBenefitsTable,
		UserWithdrawsTable,
	}
)

func init() {
}
