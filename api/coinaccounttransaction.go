// +build !codeanalysis

package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	"github.com/NpoolPlatform/cloud-hashing-billing/pkg/crud/coin-account-transaction"          //nolint
	mw "github.com/NpoolPlatform/cloud-hashing-billing/pkg/middleware/coin-account-transaction" //nolint

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCoinAccountTransaction(ctx context.Context, in *npool.CreateCoinAccountTransactionRequest) (*npool.CreateCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create coin account transaction error: %v", err)
		return &npool.CreateCoinAccountTransactionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransaction(ctx context.Context, in *npool.GetCoinAccountTransactionRequest) (*npool.GetCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction error: %v", err)
		return &npool.GetCoinAccountTransactionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByCoinAccount(ctx context.Context, in *npool.GetCoinAccountTransactionsByCoinAccountRequest) (*npool.GetCoinAccountTransactionsByCoinAccountResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByCoinAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by coin account error: %v", err)
		return &npool.GetCoinAccountTransactionsByCoinAccountResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByState(ctx context.Context, in *npool.GetCoinAccountTransactionsByStateRequest) (*npool.GetCoinAccountTransactionsByStateResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByState(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by state error: %v", err)
		return &npool.GetCoinAccountTransactionsByStateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByGoodState(ctx context.Context, in *npool.GetCoinAccountTransactionsByGoodStateRequest) (*npool.GetCoinAccountTransactionsByGoodStateResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByGoodState(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by state error: %v", err)
		return &npool.GetCoinAccountTransactionsByGoodStateResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByApp(ctx context.Context, in *npool.GetCoinAccountTransactionsByAppRequest) (*npool.GetCoinAccountTransactionsByAppResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by app error: %v", err)
		return &npool.GetCoinAccountTransactionsByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByOtherApp(ctx context.Context, in *npool.GetCoinAccountTransactionsByOtherAppRequest) (*npool.GetCoinAccountTransactionsByOtherAppResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByApp(ctx, &npool.GetCoinAccountTransactionsByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by app error: %v", err)
		return &npool.GetCoinAccountTransactionsByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetCoinAccountTransactionsByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}

func (s *Server) GetCoinAccountTransactionsByAppUser(ctx context.Context, in *npool.GetCoinAccountTransactionsByAppUserRequest) (*npool.GetCoinAccountTransactionsByAppUserResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by app user error: %v", err)
		return &npool.GetCoinAccountTransactionsByAppUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByAppUserCoin(ctx context.Context, in *npool.GetCoinAccountTransactionsByAppUserCoinRequest) (*npool.GetCoinAccountTransactionsByAppUserCoinResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByAppUserCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by app user coin error: %v", err)
		return &npool.GetCoinAccountTransactionsByAppUserCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionsByCoin(ctx context.Context, in *npool.GetCoinAccountTransactionsByCoinRequest) (*npool.GetCoinAccountTransactionsByCoinResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactionsByCoin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction by coin error: %v", err)
		return &npool.GetCoinAccountTransactionsByCoinResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactions(ctx context.Context, in *npool.GetCoinAccountTransactionsRequest) (*npool.GetCoinAccountTransactionsResponse, error) {
	resp, err := coinaccounttransaction.GetCoinAccountTransactions(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction error: %v", err)
		return &npool.GetCoinAccountTransactionsResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateCoinAccountTransaction(ctx context.Context, in *npool.UpdateCoinAccountTransactionRequest) (*npool.UpdateCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update coin account transaction error: %v", err)
		return &npool.UpdateCoinAccountTransactionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetCoinAccountTransactionDetail(ctx context.Context, in *npool.GetCoinAccountTransactionDetailRequest) (*npool.GetCoinAccountTransactionDetailResponse, error) {
	resp, err := mw.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get coin account transaction detail error: %v", err)
		return &npool.GetCoinAccountTransactionDetailResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteCoinAccountTransaction(ctx context.Context, in *npool.DeleteCoinAccountTransactionRequest) (*npool.DeleteCoinAccountTransactionResponse, error) {
	resp, err := coinaccounttransaction.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete coin account transaction error: %v", err)
		return &npool.DeleteCoinAccountTransactionResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
