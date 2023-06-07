package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"

	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithCountry(&in.Country),
		country1.WithFlag(&in.Flag),
		country1.WithCode(&in.Code),
		country1.WithShort(&in.Short),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLang",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountryResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateCountries(ctx context.Context, in *npool.CreateCountriesRequest) (*npool.CreateCountriesResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCountries",
			"In", in,
			"Error", err,
		)
		return &npool.CreateCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateCountries(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateCountries", "error", err)
		return &npool.CreateCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateCountriesResponse{
		Infos: infos,
	}, nil
}
