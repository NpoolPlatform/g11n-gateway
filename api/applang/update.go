package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"

	applangmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

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
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("UpdateLang", "AppID", in.GetAppID(), "error", err)
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := applangmgrcli.ExistLangConds(ctx, &applangmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetID(),
		},
	})
	if err != nil {
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.UpdateLangResponse{}, status.Error(codes.InvalidArgument, "AppLang not exist")
	}

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
