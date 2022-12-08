package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteCountry(ctx context.Context, in *npool.DeleteCountryRequest) (*npool.DeleteCountryResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteCountry", "ID", in.GetID(), "error", err)
		return &npool.DeleteCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: check id belong to app id

	info, err := appcountry1.DeleteCountry(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteCountry", "error", err)
		return &npool.DeleteCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCountryResponse{
		Info: info,
	}, nil
}
