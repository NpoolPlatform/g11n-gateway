package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"

	appcountrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/appcountry"
	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCountry(ctx context.Context, in *npool.DeleteCountryRequest) (*npool.DeleteCountryResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteCountry", "ID", in.GetID(), "error", err)
		return &npool.DeleteCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := appcountrymgrcli.ExistCountryConds(ctx, &appcountrymgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetID(),
		},
	})
	if err != nil {
		return &npool.DeleteCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.DeleteCountryResponse{}, status.Error(codes.InvalidArgument, "AppCountry not exist")
	}

	info, err := appcountry1.DeleteCountry(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteCountry", "error", err)
		return &npool.DeleteCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCountryResponse{
		Info: info,
	}, nil
}
