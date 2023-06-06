package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"
	langmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"
	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	handler, err := lang1.NewHandler(
		ctx,
		lang1.WithLang(&in.Lang),
		lang1.WithName(&in.Name),
		lang1.WithLogo(&in.Logo),
		lang1.WithShort(&in.Short),
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

	langs := []string{}
	for _, c := range in.GetInfos() {
		langs = append(langs, c.GetLang())
	}

	infos, _, err := langmwcli.GetLangs(ctx, &langmwpb.Conds{
		Langs: &basetypes.StringSliceVal{
			Op:    cruder.IN,
			Value: langs,
		},
	}, int32(0), int32(len(langs)))
	if err != nil {
		logger.Sugar().Errorw("CreateLangs", "error", err)
		return &npool.CreateLangsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	reqs := []*langmwpb.LangReq{}
	outs := []*langmwpb.Lang{}

	for _, info := range in.GetInfos() {
		var _info1 *langmwpb.Lang

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

	reqinfos, err := handler.CreateLangs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLangs",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLangsResponse{}, status.Error(codes.Internal, err.Error())
	}

	reqinfos = append(reqinfos, outs...)

	return &npool.CreateLangsResponse{
		Infos: reqinfos,
	}, nil
}
