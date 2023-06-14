package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithID(&in.ID),
		message1.WithAppID(&in.AppID),
		message1.WithLangID(in.TargetLangID),
		message1.WithMessageID(in.MessageID),
		message1.WithMessage(in.Message),
		message1.WithGetIndex(in.GetIndex),
		message1.WithDisabled(in.Disabled),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.UpdateMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateMessage",
			"In", in,
			"Error", err,
		)
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
		TargetLangID: in.TargetLangID,
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
