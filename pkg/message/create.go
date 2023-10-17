package message

import (
	"context"

	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (h *Handler) CreateMessage(ctx context.Context) (*npool.Message, error) {
	info, err := messagemwcli.CreateMessage(ctx, &messagemwpb.MessageReq{
		AppID:     h.AppID,
		LangID:    h.LangID,
		MessageID: h.MessageID,
		Message:   h.Message,
		GetIndex:  h.GetIndex,
		Disabled:  h.Disabled,
	})
	if err != nil {
		return nil, err
	}
	outs, err := Expand(ctx, []*messagemwpb.Message{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}

func (h *Handler) CreateMessages(ctx context.Context) ([]*npool.Message, error) {
	infos, err := messagemwcli.CreateMessages(ctx, h.Reqs)
	if err != nil {
		return nil, err
	}
	outs, err := Expand(ctx, infos)
	if err != nil {
		return nil, err
	}

	return outs, nil
}
