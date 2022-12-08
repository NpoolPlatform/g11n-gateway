package lang

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	lang1 "github.com/NpoolPlatform/g11n-gateway/pkg/lang"
	langmgrapi "github.com/NpoolPlatform/g11n-manager/api/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLang(ctx context.Context, in *npool.CreateLangRequest) (*npool.CreateLangResponse, error) {
	req := &langmgrpb.LangReq{
		ID:    in.ID,
		Lang:  &in.Lang,
		Logo:  &in.Logo,
		Name:  &in.Name,
		Short: &in.Short,
	}

	if err := langmgrapi.Validate(req); err != nil {
		logger.Sugar().Errorf("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := lang1.CreateLang(ctx, req)
	if err != nil {
		logger.Sugar().Errorf("CreateLang", "error", err)
		return &npool.CreateLangResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateLangResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateLangs(ctx context.Context, in *npool.CreateLangsRequest) (*npool.CreateLangsResponse, error) {
	return nil, nil
}
