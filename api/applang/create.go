package applang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	applang1 "github.com/NpoolPlatform/g11n-gateway/pkg/applang"

	applangmgrapi "github.com/NpoolPlatform/g11n-manager/api/applang"
	applangmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/applang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	exist, err := applangmgrcli.ExistLangConds(ctx, &applangmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("CreateLang", "error", "Lang is exist")
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, "Lang is exist")
	}

	// TODO: check app and lang exist

	req := &applangmgrpb.LangReq{
		AppID:  &in.TargetAppID,
		LangID: &in.TargetLangID,
		Main:   in.Main,
	}

	if err := applangmgrapi.Validate(req); err != nil {
		logger.Sugar().Errorw("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := applang1.CreateLang(ctx, req)
	if err != nil {
		logger.Sugar().Errorw("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}
