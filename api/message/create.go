//nolint:nolintlint,dupl
package message

import (
	"context"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithAppID(&in.AppID),
		message1.WithLangID(&in.TargetLangID),
		message1.WithMessageID(&in.MessageID),
		message1.WithMessage(&in.Message),
		message1.WithGetIndex(in.GetIndex),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.CreateMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessage",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppMessage(ctx context.Context, in *npool.CreateAppMessageRequest) (*npool.CreateAppMessageResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithAppID(&in.TargetAppID),
		message1.WithLangID(&in.TargetLangID),
		message1.WithMessageID(&in.MessageID),
		message1.WithMessage(&in.Message),
		message1.WithGetIndex(in.GetIndex),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppMessage",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppMessage",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessages",
			"In", in,
			"Error", err,
		)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, err := handler.CreateMessages(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessagesResponse{
		Infos: infos,
	}, nil
}

func (s *Server) CreateAppMessages(ctx context.Context, in *npool.CreateAppMessagesRequest) (*npool.CreateAppMessagesResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithAppID(&in.TargetAppID),
		message1.WithLangID(&in.TargetLangID),
		message1.WithAppReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateMessages",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	outs, err := handler.CreateMessages(ctx)
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateAppMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAppMessagesResponse{
		Infos: outs,
	}, nil
}
