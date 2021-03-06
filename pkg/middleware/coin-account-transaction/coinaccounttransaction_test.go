package coinaccounttransaction

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/test-init" //nolint
	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-info"        //nolint
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-transaction" //nolint

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

func TestGet(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinTypeID := uuid.New().String()

	account1 := npool.CoinAccountInfo{
		CoinTypeID: coinTypeID,
		Address:    uuid.New().String(),
	}
	account2 := npool.CoinAccountInfo{
		CoinTypeID: coinTypeID,
		Address:    uuid.New().String(),
	}

	resp1, err := coinaccountinfo.Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &account1,
	})
	assert.Nil(t, err)

	resp2, err := coinaccountinfo.Create(context.Background(), &npool.CreateCoinAccountRequest{
		Info: &account2,
	})
	assert.Nil(t, err)

	transaction := npool.CoinAccountTransaction{
		AppID:              uuid.New().String(),
		UserID:             uuid.New().String(),
		GoodID:             uuid.New().String(),
		FromAddressID:      resp1.Info.ID,
		ToAddressID:        resp2.Info.ID,
		CoinTypeID:         coinTypeID,
		Amount:             0.13,
		Message:            "test transaction",
		ChainTransactionID: "adfjklasjfdlksajf",
	}
	resp3, err := coinaccounttransaction.Create(context.Background(), &npool.CreateCoinAccountTransactionRequest{
		Info: &transaction,
	})
	assert.Nil(t, err)

	resp4, err := Get(context.Background(), &npool.GetCoinAccountTransactionDetailRequest{
		ID: resp3.Info.ID,
	})
	if assert.Nil(t, err) {
		assert.Equal(t, resp4.Detail.ID, resp3.Info.ID)
		assert.Equal(t, resp4.Detail.UserID, resp3.Info.UserID)
		assert.Equal(t, resp4.Detail.AppID, resp3.Info.AppID)

		assert.Equal(t, resp4.Detail.FromAddress.ID, resp1.Info.ID)
		assert.Equal(t, resp4.Detail.FromAddress.CoinTypeID, resp1.Info.CoinTypeID)
		assert.Equal(t, resp4.Detail.FromAddress.Address, resp1.Info.Address)

		assert.Equal(t, resp4.Detail.ToAddress.ID, resp2.Info.ID)
		assert.Equal(t, resp4.Detail.ToAddress.CoinTypeID, resp2.Info.CoinTypeID)
		assert.Equal(t, resp4.Detail.ToAddress.Address, resp2.Info.Address)

		assert.Equal(t, resp4.Detail.Amount, resp3.Info.Amount)
		assert.Equal(t, resp4.Detail.Message, resp3.Info.Message)
		assert.Equal(t, resp4.Detail.ChainTransactionID, resp3.Info.ChainTransactionID)
		assert.Equal(t, resp4.Detail.State, resp3.Info.State)
	}
}
