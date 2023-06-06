package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	hangler, err := message1.NewHandler(
		ctx,
		message1.WithAppID(&in.AppID),
		message1.WithDisabled(in.Disabled),
		message1.WithLangID(in.LangID),
		message1.WithOffset(in.GetOffset()),
		message1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessages",
			"In", in,
			"Error", err,
		)
		return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := hangler.GetMessages(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetMessages",
			"In", in,
			"Error", err,
		)
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
