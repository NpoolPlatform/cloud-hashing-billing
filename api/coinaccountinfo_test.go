package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"
)

func assertCoinAccount(t *testing.T, actual, expected *npool.CoinAccountInfo) {
	assert.Equal(t, actual.CoinTypeID, expected.CoinTypeID)
	assert.Equal(t, actual.Address, expected.Address)
	assert.Equal(t, actual.PlatformHoldPrivateKey, expected.PlatformHoldPrivateKey)
}

func TestCoinAccountInfoCRUD(t *testing.T) { //nolint
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	coinAccount := npool.CoinAccountInfo{
		CoinTypeID: uuid.New().String(),
		Address:    uuid.New().String(),
	}
	firstCreateInfo := npool.CreateCoinAccountResponse{}

	cli := resty.New()

	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.CreateCoinAccountRequest{
			Info: &coinAccount,
		}).
		Post("http://localhost:50030/v1/create/coin/account")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		err := json.Unmarshal(resp.Body(), &firstCreateInfo)
		if assert.Nil(t, err) {
			assert.NotEqual(t, firstCreateInfo.Info.ID, uuid.UUID{}.String())
			assertCoinAccount(t, firstCreateInfo.Info, &coinAccount)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinAccountRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:50030/v1/get/coin/account")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetCoinAccountResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertCoinAccount(t, info.Info, &coinAccount)
		}
	}

	resp, err = cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(npool.GetCoinAccountRequest{
			ID: firstCreateInfo.Info.ID,
		}).
		Post("http://localhost:50030/v1/delete/coin/account")
	if assert.Nil(t, err) {
		assert.Equal(t, 200, resp.StatusCode())
		info := npool.GetCoinAccountResponse{}
		err := json.Unmarshal(resp.Body(), &info)
		if assert.Nil(t, err) {
			assert.Equal(t, info.Info.ID, firstCreateInfo.Info.ID)
			assertCoinAccount(t, info.Info, &coinAccount)
		}
	}
}
