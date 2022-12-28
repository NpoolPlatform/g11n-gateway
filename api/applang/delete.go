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

func (s *Server) DeleteLang(ctx context.Context, in *npool.DeleteLangRequest) (*npool.DeleteLangResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteLang", "ID", in.GetID(), "error", err)
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := applangmgrcli.ExistLangConds(ctx, &applangmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		ID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetID(),
		},
	})
	if err != nil {
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.DeleteLangResponse{}, status.Error(codes.InvalidArgument, "AppLang not exist")
	}

	info, err := applang1.DeleteLang(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteLang", "error", err)
		return &npool.DeleteLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteLangResponse{
		Info: info,
	}, nil
}
