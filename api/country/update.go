package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCountry(ctx context.Context, in *npool.UpdateCountryRequest) (*npool.UpdateCountryResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithID(&in.ID),
		country1.WithCountry(in.Country),
		country1.WithCode(in.Code),
		country1.WithFlag(in.Flag),
		country1.WithShort(in.Short),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCountry",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCountryResponse{
		Info: info,
	}, nil
}
