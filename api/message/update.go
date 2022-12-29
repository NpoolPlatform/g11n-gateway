package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	messagemgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	applangmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateMessage", "ID", in.GetID(), "error", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("UpdateMessage", "AppID", in.GetAppID(), "error", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetMessageID() == "" {
		logger.Sugar().Errorw("UpdateMessage", "MessageID", in.GetMessageID())
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "MessageID is invalid")
	}
	if in.GetMessage() == "" {
		logger.Sugar().Errorw("UpdateMessage", "Message", in.GetMessage())
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "Message is invalid")
	}

	exist, err := messagemgrcli.ExistMessageConds(ctx, &messagemgrpb.Conds{
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
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "Message not exist")
	}

	exist, err = applangmgrcli.ExistLangConds(ctx, &applangmgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
	})
	if err != nil {
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "AppLang not exist")
	}

	if in.MessageID != nil {
		exist, err := messagemgrcli.ExistMessageConds(ctx, &messagemgrpb.Conds{
			AppID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetAppID(),
			},
			LangID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetTargetLangID(),
			},
			ID: &commonpb.StringVal{
				Op:    cruder.NEQ,
				Value: in.GetID(),
			},
			MessageID: &commonpb.StringVal{
				Op:    cruder.EQ,
				Value: in.GetMessageID(),
			},
		})
		if err != nil {
			return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		if exist {
			return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "MessageID exist")
		}
	}

	info, err := message1.UpdateMessage(ctx, &messagemgrpb.MessageReq{
		ID:        &in.ID,
		MessageID: in.MessageID,
		Message:   in.Message,
		GetIndex:  in.GetIndex,
		Disabled:  in.Disabled,
	})
	if err != nil {
		logger.Sugar().Errorw("UpdateMessage", "error", err)
		return &npool.UpdateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateAppMessage(ctx context.Context, in *npool.UpdateAppMessageRequest) (*npool.UpdateAppMessageResponse, error) {
	r, err := s.UpdateMessage(ctx, &npool.UpdateMessageRequest{
		ID:           in.GetID(),
		AppID:        in.GetTargetAppID(),
		TargetLangID: in.GetTargetLangID(),
		MessageID:    in.MessageID,
		Message:      in.Message,
		GetIndex:     in.GetIndex,
		Disabled:     in.Disabled,
	})
	if err != nil {
		return &npool.UpdateAppMessageResponse{}, err
	}

	return &npool.UpdateAppMessageResponse{
		Info: r.Info,
	}, nil
}
