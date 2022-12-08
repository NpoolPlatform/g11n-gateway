package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	// TODO: check id belong to app id

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("UpdateMessage", "ID", in.GetID(), "error", err)
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
		ID:        in.ID,
		AppID:     in.TargetAppID,
		MessageID: in.MessageID,
		Message:   in.Message,
		GetIndex:  in.GetIndex,
		Disabled:  in.Disabled,
	})
	if err != nil {
		return &npool.UpdateAppMessageResponse{}, err
	}

	return &npool.UpdateAppMessageResponse{
		Info: r.Info,
	}, nil
}
