package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCountry(ctx context.Context, in *npool.DeleteCountryRequest) (*npool.DeleteCountryResponse, error) {
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithID(&in.ID),
		appcountry1.WithAppID(&in.TargetAppID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCountry",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteCountry(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCountry",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCountryResponse{
		Info: info,
	}, nil
}
