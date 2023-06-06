package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteMessage(ctx context.Context, in *npool.DeleteMessageRequest) (*npool.DeleteMessageResponse, error) {
	handler, err := message1.NewHandler(
		ctx,
		message1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteAppMessage(ctx context.Context, in *npool.DeleteAppMessageRequest) (*npool.DeleteAppMessageResponse, error) {
	req := &messagemwpb.MessageReq{}
	req.ID = &in.ID
	req.AppID = &in.TargetAppID
	handler, err := message1.NewHandler(
		ctx,
		message1.WithID(req.ID),
		message1.WithAppID(req.AppID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := handler.DeleteMessage(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteMessage",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAppMessageResponse{
		Info: info,
	}, nil
}
