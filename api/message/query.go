package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := message1.GetMessages(ctx, &messagemgrpb.Conds{}, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetMessages", "error", err)
		return &npool.GetMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMessagesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
