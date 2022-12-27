package message

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	message1 "github.com/NpoolPlatform/g11n-gateway/pkg/message1"
	messagemgrapi "github.com/NpoolPlatform/g11n-manager/api/message"
	messagemgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/message"
	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	langmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/lang"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	exist, err := messagemgrcli.ExistMessageConds(ctx, &messagemgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
		MessageID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetMessageID(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateMessage", "error", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if exist {
		logger.Sugar().Errorw("CreateMessage", "error", "Message is exist")
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, "Message is exist")
	}

	exist, err = langmgrcli.ExistLang(ctx, in.GetTargetLangID())
	if err != nil {
		logger.Sugar().Errorw("CreateMessage", "error", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		logger.Sugar().Errorw("CreateMessage", "error", "Lang isn't exist")
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, "Lang isn't exist")
	}

	app, err := appmwcli.GetApp(ctx, in.GetAppID())
	if err != nil {
		logger.Sugar().Errorw("CreateMessage", "error", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if app == nil {
		logger.Sugar().Errorw("CreateMessage", "error", "App isn't exist")
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, "App isn't exist")
	}

	req := &messagemgrpb.MessageReq{
		AppID:     &in.AppID,
		LangID:    &in.TargetLangID,
		MessageID: &in.MessageID,
		Message:   &in.Message,
		GetIndex:  in.GetIndex,
	}

	if err := messagemgrapi.Validate(req); err != nil {
		logger.Sugar().Errorw("CreateMessage", "error", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := message1.CreateMessage(ctx, req)
	if err != nil {
		logger.Sugar().Errorw("CreateMessage", "error", err)
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessageResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateAppMessage(ctx context.Context, in *npool.CreateAppMessageRequest) (*npool.CreateAppMessageResponse, error) {
	r, err := s.CreateMessage(ctx, &npool.CreateMessageRequest{
		AppID:        in.TargetAppID,
		TargetLangID: in.TargetLangID,
		MessageID:    in.MessageID,
		Message:      in.Message,
		GetIndex:     in.GetIndex,
	})
	if err != nil {
		return &npool.CreateAppMessageResponse{}, err
	}

	return &npool.CreateAppMessageResponse{
		Info: r.Info,
	}, nil
}

//nolint:funlen,gocyclo
func (s *Server) CreateMessages(
	ctx context.Context,
	in *npool.CreateMessagesRequest,
) (
	*npool.CreateMessagesResponse,
	error,
) {
	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateMessages", "error", "Infos is empty")
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	if err := messagemgrapi.Duplicate(in.GetInfos()); err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	for _, info := range in.GetInfos() {
		if info.GetAppID() != in.GetAppID() || info.GetLangID() != in.GetTargetLangID() {
			return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, "Infos is invalid")
		}
	}

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("CreateMessages", "AppID", in.GetAppID(), "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetTargetLangID()); err != nil {
		logger.Sugar().Errorw("CreateMessages", "TargetLangID", in.GetTargetLangID(), "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	msgIDs := []string{}
	for _, info := range in.GetInfos() {
		msgIDs = append(msgIDs, info.GetMessageID())
	}

	infos, _, err := messagemwcli.GetMessages(ctx, &messagemgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetAppID(),
		},
		LangID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetTargetLangID(),
		},
		MessageIDs: &commonpb.StringSliceVal{
			Op:    cruder.IN,
			Value: msgIDs,
		},
	}, int32(0), int32(len(msgIDs)))
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	msgs := []*messagemgrpb.MessageReq{}
	for _, info := range in.GetInfos() {
		exist := false
		for _, info1 := range infos {
			if info.GetMessageID() == info1.MessageID {
				exist = true
				break
			}
		}
		if !exist {
			msgs = append(msgs, info)
		}
	}

	if len(msgs) == 0 {
		outs, err := message1.Expand(ctx, infos)
		if err != nil {
			logger.Sugar().Errorw("CreateMessages", "error", err)
			return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, err.Error())
		}
		return &npool.CreateMessagesResponse{
			Infos: outs,
		}, nil
	}

	exist, err := langmgrcli.ExistLang(ctx, in.GetTargetLangID())
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if !exist {
		logger.Sugar().Errorw("CreateMessages", "error", "Lang isn't exist")
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, "Lang isn't exist")
	}

	app, err := appmwcli.GetApp(ctx, in.GetAppID())
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if app == nil {
		logger.Sugar().Errorw("CreateMessage", "error", "App isn't exist")
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, "App isn't exist")
	}

	outs, err := message1.CreateMessages(ctx, msgs)
	if err != nil {
		logger.Sugar().Errorw("CreateMessages", "error", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessagesResponse{
		Infos: outs,
	}, nil
}

func (s *Server) CreateAppMessages(
	ctx context.Context,
	in *npool.CreateAppMessagesRequest,
) (
	*npool.CreateAppMessagesResponse,
	error,
) {
	infos := []*messagemgrpb.MessageReq{}
	for _, info := range in.GetInfos() {
		info.AppID = &in.TargetAppID
		info.LangID = &in.TargetLangID
		infos = append(infos, info)
	}

	r, err := s.CreateMessages(ctx, &npool.CreateMessagesRequest{
		AppID:        in.GetTargetAppID(),
		TargetLangID: in.GetTargetLangID(),
		Infos:        infos,
	})
	if err != nil {
		return &npool.CreateAppMessagesResponse{}, err
	}

	return &npool.CreateAppMessagesResponse{
		Infos: r.Infos,
	}, nil
}
