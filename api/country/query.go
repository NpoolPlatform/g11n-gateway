package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"

	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	handler, err := country1.NewHandler(
		ctx,
		country1.WithOffset(in.GetOffset()),
		country1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountries",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetCountries(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCountries",
			"In", in,
			"Error", err,
		)
		return &npool.GetCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
