package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetMessages", "AppID", in.GetAppID(), "error", err)
		return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.LangID != nil {
		if _, err := uuid.Parse(in.GetLangID()); err != nil {
			logger.Sugar().Errorw("GetMessages", "LangID", in.GetLangID(), "error", err)
			return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	conds := &messagemgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		Disabled: &commonpb.BoolVal{
			Op:    cruder.EQ,
			Value: in.GetDisabled(),
		},
	}
	if in.LangID != nil {
		conds.LangID = &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetLangID(),
		}
	}

	infos, total, err := message1.GetMessages(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetMessages", "error", err)
		return &npool.GetMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMessagesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetAppMessages(ctx context.Context, in *npool.GetAppMessagesRequest) (*npool.GetAppMessagesResponse, error) {
	r, err := s.GetMessages(ctx, &npool.GetMessagesRequest{
		AppID:    in.TargetAppID,
		Disabled: in.Disabled,
		Offset:   in.Offset,
		Limit:    in.Limit,
	})
	if err != nil {
		return &npool.GetAppMessagesResponse{}, err
	}

	return &npool.GetAppMessagesResponse{
		Infos: r.Infos,
		Total: r.Total,
	}, nil
}
