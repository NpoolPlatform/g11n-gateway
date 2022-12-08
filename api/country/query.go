package country

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/country"
	countrymgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	country1 "github.com/NpoolPlatform/g11n-gateway/pkg/country"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetCountries(ctx context.Context, in *npool.GetCountriesRequest) (*npool.GetCountriesResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := country1.GetCountries(ctx, &countrymgrpb.Conds{}, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetCountries", "error", err)
		return &npool.GetCountriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCountriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
