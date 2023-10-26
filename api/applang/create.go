//nolint:nolintlint,dupl
package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithLangID(&in.TargetLangID, true),
		applang1.WithAppID(&in.TargetAppID, true),
		applang1.WithMain(in.Main, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLang",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLang",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}
