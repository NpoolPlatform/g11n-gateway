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

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	handler, err := applang1.NewHandler(
		ctx,
		applang1.WithID(&in.ID, true),
		applang1.WithAppID(&in.AppID, true),
		applang1.WithMain(in.Main, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLang",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateLang(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateLang",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateLangResponse{
		Info: info,
	}, nil
}
