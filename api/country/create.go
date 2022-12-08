package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"
	countrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"
	countrymgrapi "github.com/NpoolPlatform/g11n-manager/api/country"
	countrymgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	exist, err := countrymgrcli.ExistCountryConds(ctx, &countrymgrpb.Conds{
		Country: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.Country,
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

	req := &countrymgrpb.CountryReq{
		ID:      in.ID,
		Country: &in.Country,
		Flag:    &in.Flag,
		Code:    &in.Code,
		Short:   &in.Short,
	}

	if err := countrymgrapi.Validate(req); err != nil {
		logger.Sugar().Errorw("CreateCountry", "error", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := country1.CreateCountry(ctx, req)
	if err != nil {
		logger.Sugar().Errorw("CreateCountry", "error", err)
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountryResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	if err := countrymgrapi.Duplicate(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateCountries", "error", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := country1.CreateCountries(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateCountries", "error", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountriesResponse{
		Infos: infos,
	}, nil
}
