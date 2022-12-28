package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"
	appcountrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"

	appcountrymgrapi "github.com/NpoolPlatform/g11n-manager/api/appcountry"
	appcountrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/appcountry"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	countrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	exist, err := appcountrymgrcli.ExistCountryConds(ctx, &appcountrymgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		CountryID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetCountryID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateCountry", "error", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("CreateCountry", "error", "Country is exist")
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, "Country is exist")
	}

	// TODO: check app and lang exist
	app, err := appmwcli.GetApp(ctx, in.GetTargetAppID())
	if err != nil {
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if app == nil {
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, "App not exist")
	}

	exist, err = countrymgrcli.ExistCountry(ctx, in.GetCountryID())
	if err != nil {
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, "Country not exist")
	}

	req := &appcountrymgrpb.CountryReq{
		AppID:     &in.TargetAppID,
		CountryID: &in.CountryID,
	}

	if err := appcountrymgrapi.Validate(req); err != nil {
		logger.Sugar().Errorw("CreateCountry", "error", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := appcountry1.CreateCountry(ctx, req)
	if err != nil {
		logger.Sugar().Errorw("CreateCountry", "error", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountryResponse{
		Info: info,
	}, nil
}
