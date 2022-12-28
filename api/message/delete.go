package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	messagemgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteMessage(ctx context.Context, in *npool.DeleteMessageRequest) (*npool.DeleteMessageResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("DeleteMessage", "ID", in.GetID(), "error", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("DeleteMessage", "AppID", in.GetAppID(), "error", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
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
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, "Message not exist")
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
		ID:    in.GetID(),
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		return &npool.DeleteAppMessageResponse{}, err
	}

	return &npool.DeleteAppMessageResponse{
		Info: r.Info,
	}, nil
}
