package platformbenefit

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/db/ent/platformbenefit"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/google/uuid"

	"golang.org/x/xerrors"
)

func validatePlatformBenefit(info *npool.PlatformBenefit) error {
	if _, err := uuid.Parse(info.GetGoodID()); err != nil {
		return xerrors.Errorf("invalid good id: %v", err)
	}
	if _, err := uuid.Parse(info.GetBenefitAccountID()); err != nil {
		return xerrors.Errorf("invalid benefit account id: %v", err)
	}
	if info.GetLastBenefitTimestamp() == 0 {
		return xerrors.Errorf("invalid last benefit stamp")
	}
	return nil
}

func dbRowToPlatformBenefit(row *ent.PlatformBenefit) *npool.PlatformBenefit {
	return &npool.PlatformBenefit{
		ID:                   row.ID.String(),
		GoodID:               row.GoodID.String(),
		BenefitAccountID:     row.BenefitAccountID.String(),
		Amount:               price.DBPriceToVisualPrice(row.Amount),
		ChainTransactionID:   row.ChainTransactionID,
		CreateAt:             row.CreateAt,
		LastBenefitTimestamp: row.LastBenefitTimestamp,
	}
}

func Create(ctx context.Context, in *npool.CreatePlatformBenefitRequest) (*npool.CreatePlatformBenefitResponse, error) {
	if err := validatePlatformBenefit(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	info, err := cli.
		PlatformBenefit.
		Create().
		SetGoodID(uuid.MustParse(in.GetInfo().GetGoodID())).
		SetBenefitAccountID(uuid.MustParse(in.GetInfo().GetBenefitAccountID())).
		SetAmount(price.VisualPriceToDBPrice(in.GetInfo().GetAmount())).
		SetChainTransactionID(in.GetInfo().GetChainTransactionID()).
		SetLastBenefitTimestamp(in.GetInfo().GetLastBenefitTimestamp()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create platform benefit: %v", err)
	}

	return &npool.CreatePlatformBenefitResponse{
		Info: dbRowToPlatformBenefit(info),
	}, nil
}

func GetLatestByGood(ctx context.Context, in *npool.GetLatestPlatformBenefitByGoodRequest) (*npool.GetLatestPlatformBenefitByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		PlatformBenefit.
		Query().
		Order(
			ent.Desc(platformbenefit.FieldCreateAt),
		).
		Where(
			platformbenefit.And(
				platformbenefit.GoodID(goodID),
			),
		).
		Limit(1).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query platform benefit: %v", err)
	}

	var benefit *npool.PlatformBenefit
	for _, info := range infos {
		benefit = dbRowToPlatformBenefit(info)
		break
	}

	return &npool.GetLatestPlatformBenefitByGoodResponse{
		Info: benefit,
	}, nil
}

func GetByGood(ctx context.Context, in *npool.GetPlatformBenefitsByGoodRequest) (*npool.GetPlatformBenefitsByGoodResponse, error) {
	goodID, err := uuid.Parse(in.GetGoodID())
	if err != nil {
		return nil, xerrors.Errorf("invalid good id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		PlatformBenefit.
		Query().
		Where(
			platformbenefit.And(
				platformbenefit.GoodID(goodID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query platform benefit: %v", err)
	}

	benefits := []*npool.PlatformBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToPlatformBenefit(info))
	}

	return &npool.GetPlatformBenefitsByGoodResponse{
		Infos: benefits,
	}, nil
}

func Get(ctx context.Context, in *npool.GetPlatformBenefitRequest) (*npool.GetPlatformBenefitResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		PlatformBenefit.
		Query().
		Where(
			platformbenefit.And(
				platformbenefit.ID(id),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query platform benefit: %v", err)
	}

	var benefit *npool.PlatformBenefit
	for _, info := range infos {
		benefit = dbRowToPlatformBenefit(info)
		break
	}

	return &npool.GetPlatformBenefitResponse{
		Info: benefit,
	}, nil
}

func GetAll(ctx context.Context, in *npool.GetPlatformBenefitsRequest) (*npool.GetPlatformBenefitsResponse, error) {
	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db client: %v", err)
	}

	infos, err := cli.
		PlatformBenefit.
		Query().
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query platform benefit: %v", err)
	}

	benefits := []*npool.PlatformBenefit{}
	for _, info := range infos {
		benefits = append(benefits, dbRowToPlatformBenefit(info))
	}

	return &npool.GetPlatformBenefitsResponse{
		Infos: benefits,
	}, nil
}
