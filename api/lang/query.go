package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetLangs(ctx context.Context, in *npool.GetLangsRequest) (*npool.GetLangsResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := lang1.GetLangs(ctx, &langmgrpb.Conds{}, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetLangs", "error", err)
		return &npool.GetLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetLangsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
