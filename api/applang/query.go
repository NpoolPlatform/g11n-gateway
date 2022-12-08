package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"
	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetLangs", "AppID", in.GetAppID(), "error", err)
		return &npool.GetLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := applang1.GetLangs(ctx, in.GetAppID(), in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetLangs", "error", err)
		return &npool.GetLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLangsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppLangs(ctx context.Context, in *npool.GetAppLangsRequest) (*npool.GetAppLangsResponse, error) {
	r, err := s.GetLangs(ctx, &npool.GetLangsRequest{
		AppID: in.TargetAppID,
	})
	if err != nil {
		return &npool.GetAppLangsResponse{}, err
	}

	return &npool.GetAppLangsResponse{
		Infos: r.Infos,
		Total: r.Total,
	}, nil
}
