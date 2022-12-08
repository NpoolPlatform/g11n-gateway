package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateLang", "ID", in.GetID(), "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetLang() == "" {
		logger.Sugar().Errorw("UpdateLang", "Lang", in.GetLang())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Lang is invalid")
	}
	if in.GetLogo() == "" {
		logger.Sugar().Errorw("UpdateLang", "Logo", in.GetLogo())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Logo is invalid")
	}
	if in.GetName() == "" {
		logger.Sugar().Errorw("UpdateLang", "Name", in.GetName())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Name is invalid")
	}
	if in.GetShort() == "" {
		logger.Sugar().Errorw("UpdateLang", "Short", in.GetShort())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Short is invalid")
	}

	info, err := lang1.UpdateLang(ctx, &langmgrpb.LangReq{
		ID:    &in.ID,
		Lang:  in.Lang,
		Logo:  in.Logo,
		Name:  in.Name,
		Short: in.Short,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateLang", "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateLangResponse{
		Info: info,
	}, nil
}
