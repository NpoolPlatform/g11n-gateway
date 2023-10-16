package message

import (
	"context"

	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

//nolint:dupl
func (h *Handler) UpdateMessage(ctx context.Context) (*npool.Message, error) {
	info, err := messagemwcli.UpdateMessage(ctx, &messagemwpb.MessageReq{
		ID:        h.ID,
		AppID:     h.AppID,
		LangID:    h.LangID,
		MessageID: h.MessageID,
		Message:   h.Message,
		Disabled:  h.Disabled,
		GetIndex:  h.GetIndex,
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
