package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"

	langmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateLang(ctx context.Context, in *npool.UpdateLangRequest) (*npool.UpdateLangResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateLang", "ID", in.GetID(), "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Lang != nil && in.GetLang() == "" {
		logger.Sugar().Errorw("UpdateLang", "Lang", in.GetLang())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Lang is invalid")
	}
	if in.Logo != nil && in.GetLogo() == "" {
		logger.Sugar().Errorw("UpdateLang", "Logo", in.GetLogo())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Logo is invalid")
	}
	if in.Name != nil && in.GetName() == "" {
		logger.Sugar().Errorw("UpdateLang", "Name", in.GetName())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Name is invalid")
	}
	if in.Short != nil && in.GetShort() == "" {
		logger.Sugar().Errorw("UpdateLang", "Short", in.GetShort())
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Short is invalid")
	}

	exist, err := langmgrcli.ExistLangConds(ctx, &langmgrpb.Conds{
		ID: &commonpb.StringVal{
			Op:    cruder.NEQ,
			Value: in.GetID(),
		},
		Lang: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetLang(),
		},
	})
	if err != nil {
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if exist {
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "Lang is exist")
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
