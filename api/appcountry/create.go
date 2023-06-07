package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateCountry(ctx context.Context, in *npool.CreateCountryRequest) (*npool.CreateCountryResponse, error) {
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithCountryID(&in.CountryID),
		appcountry1.WithAppID(&in.TargetAppID),
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
