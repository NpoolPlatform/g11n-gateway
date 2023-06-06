package appcountry

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/appcountry"

	appcountry1 "github.com/NpoolPlatform/g11n-gateway/pkg/appcountry"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	handler, err := appcountry1.NewHandler(
		ctx,
		appcountry1.WithID(&in.AppID),
		appcountry1.WithOffset(in.GetOffset()),
		appcountry1.WithLimit(in.GetLimit()),
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
