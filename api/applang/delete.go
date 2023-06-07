//nolint:nolintlint,dupl
package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteLang(ctx context.Context, in *npool.DeleteLangRequest) (*npool.DeleteLangResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithID(&in.ID),
		applang1.WithAppID(&in.TargetAppID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLang",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteLang",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteLangResponse{
		Info: info,
	}, nil
}
