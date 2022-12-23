package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"
	countrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateCountry(ctx context.Context, in *npool.UpdateCountryRequest) (*npool.UpdateCountryResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateCountry", "ID", in.GetID(), "error", err)
		return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Country != nil {
		if in.GetCountry() == "" {
			logger.Sugar().Errorw("UpdateCountry", "Country", in.GetCountry())
			return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Country is invalid")
		}
	}
	if in.Flag != nil {
		if in.GetFlag() == "" {
			logger.Sugar().Errorw("UpdateCountry", "Flag", in.GetFlag())
			return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Flag is invalid")
		}
	}
	if in.Code != nil {
		if in.GetCode() == "" {
			logger.Sugar().Errorw("UpdateCountry", "Code", in.GetCode())
			return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Code is invalid")
		}
	}
	if in.Short != nil {
		if in.GetShort() == "" {
			logger.Sugar().Errorw("UpdateCountry", "Short", in.GetShort())
			return &npool.UpdateCountryResponse{}, status.Error(codes.InvalidArgument, "Short is invalid")
		}
	}

	info, err := country1.UpdateCountry(ctx, &countrymgrpb.CountryReq{
		ID:      &in.ID,
		Country: in.Country,
		Flag:    in.Flag,
		Code:    in.Code,
		Short:   in.Short,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateCountry", "error", err)
		return &npool.UpdateCountryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCountryResponse{
		Info: info,
	}, nil
}
