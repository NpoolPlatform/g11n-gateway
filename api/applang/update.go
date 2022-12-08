package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateLang", "ID", in.GetID(), "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: check id belong to app id

	info, err := applang1.UpdateLang(ctx, &applangmgrpb.LangReq{
		ID:   &in.ID,
		Main: in.Main,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateLang", "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateLangResponse{
		Info: info,
	}, nil
}