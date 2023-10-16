package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithID(&in.ID, true),
		lang1.WithLang(in.Lang, false),
		lang1.WithLogo(in.Logo, false),
		lang1.WithName(in.Name, false),
		lang1.WithShort(in.Short, false),
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
