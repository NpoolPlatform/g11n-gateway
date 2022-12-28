package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"
	langmgrapi "github.com/NpoolPlatform/g11n-manager/api/lang"
	langmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	exist, err := langmgrcli.ExistLangConds(ctx, &langmgrpb.Conds{
		Lang: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.Lang,
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

	req := &langmgrpb.LangReq{
		ID:    in.ID,
		Lang:  &in.Lang,
		Logo:  &in.Logo,
		Name:  &in.Name,
		Short: &in.Short,
	}

	if err := langmgrapi.Validate(req); err != nil {
		logger.Sugar().Errorw("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := lang1.CreateLang(ctx, req)
	if err != nil {
		logger.Sugar().Errorw("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateLangs(ctx context.Context, in *npool.CreateLangsRequest) (*npool.CreateLangsResponse, error) {
	if err := langmgrapi.Duplicate(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateLangs", "error", err)
		return &npool.CreateLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	langs := []string{}
	for _, c := range in.GetInfos() {
		langs = append(langs, c.GetLang())
	}

	infos, _, err := langmgrcli.GetLangs(ctx, &langmgrpb.Conds{
		Langs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: langs,
		},
	}, int32(0), int32(len(langs)))
	if err != nil {
		logger.Sugar().Errorw("CreateLangs", "error", err)
		return &npool.CreateLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	reqs := []*langmgrpb.LangReq{}
	outs := []*langmgrpb.Lang{}

	for _, info := range in.GetInfos() {
		var _info1 *langmgrpb.Lang

		exist := false
		for _, info1 := range infos {
			if info.GetLang() == info1.Lang {
				_info1 = info1
				exist = true
				break
			}
		}
		if !exist {
			reqs = append(reqs, info)
			continue
		}

		outs = append(outs, _info1)
	}

	if len(reqs) == 0 {
		return &npool.CreateLangsResponse{
			Infos: infos,
		}, nil
	}

	infos, err = lang1.CreateLangs(ctx, reqs)
	if err != nil {
		logger.Sugar().Errorw("CreateLangs", "error", err)
		return &npool.CreateLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	infos = append(infos, outs...)

	return &npool.CreateLangsResponse{
		Infos: infos,
	}, nil
}
