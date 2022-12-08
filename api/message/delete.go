package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteMessage(ctx context.Context, in *npool.DeleteMessageRequest) (*npool.DeleteMessageResponse, error) {
	// TODO: check id belong to app id

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteMessage", "ID", in.GetID(), "error", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := message1.DeleteMessage(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteMessage", "error", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteAppMessage(ctx context.Context, in *npool.DeleteAppMessageRequest) (*npool.DeleteAppMessageResponse, error) {
	r, err := s.DeleteMessage(ctx, &npool.DeleteMessageRequest{
		ID: in.ID,
	})
	if err != nil {
		return &npool.DeleteAppMessageResponse{}, err
	}

	return &npool.DeleteAppMessageResponse{
		Info: r.Info,
	}, nil
}
