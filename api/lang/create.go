package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithEntID(in.EntID, false),
		lang1.WithLang(&in.Lang, true),
		lang1.WithName(&in.Name, true),
		lang1.WithLogo(&in.Logo, true),
		lang1.WithShort(&in.Short, true),
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
			"CreateCoin",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateLangs(ctx context.Context, in *npool.CreateLangsRequest) (*npool.CreateLangsResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLangs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateLangs(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateLangs", "error", err)
		return &npool.CreateLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangsResponse{
		Infos: infos,
	}, nil
}
