package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithOffset(in.GetOffset()),
		lang1.WithLimit(in.GetLimit()),
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
