package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"
	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetCountries", "AppID", in.GetAppID(), "error", err)
		return &npool.GetCountriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := appcountry1.GetCountries(ctx, in.GetAppID(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCountries", "error", err)
		return &npool.GetCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppCountries(ctx context.Context, in *npool.GetAppCountriesRequest) (*npool.GetAppCountriesResponse, error) {
	r, err := s.GetCountries(ctx, &npool.GetCountriesRequest{
		AppID:  in.TargetAppID,
		Offset: in.GetOffset(),
		Limit:  in.GetLimit(),
	})
	if err != nil {
		return &npool.GetAppCountriesResponse{}, err
	}

	return &npool.GetAppCountriesResponse{
		Infos: r.Infos,
		Total: r.Total,
	}, nil
}
