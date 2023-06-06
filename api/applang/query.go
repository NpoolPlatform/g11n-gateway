package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithID(&in.AppID),
		applang1.WithOffset(in.GetOffset()),
		applang1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLangs",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetLangs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetLangs",
			"In", in,
			"Error", err,
		)
		return &npool.GetLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLangsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppLangs(ctx context.Context, in *npool.GetAppLangsRequest) (*npool.GetAppLangsResponse, error) {
	r, err := s.GetLangs(ctx, &npool.GetLangsRequest{
		AppID:  in.TargetAppID,
		Offset: in.GetOffset(),
		Limit:  in.GetLimit(),
	})
	if err != nil {
		return &npool.GetAppLangsResponse{}, err
	}

	return &npool.GetAppLangsResponse{
		Infos: r.Infos,
		Total: r.Total,
	}, nil
}
