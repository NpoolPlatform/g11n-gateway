package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteLang(ctx context.Context, in *npool.DeleteLangRequest) (*npool.DeleteLangResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteLang", "ID", in.GetID(), "error", err)
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: check id belong to app id

	info, err := applang1.DeleteLang(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteLang", "error", err)
		return &npool.DeleteLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteLangResponse{
		Info: info,
	}, nil
}
