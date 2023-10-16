package message

import (
	"context"
	"fmt"

	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (h *Handler) DeleteMessage(ctx context.Context) (*npool.Message, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	if h.AppID == nil {
		return nil, fmt.Errorf("invalid appid")
	}
	info, err := messagemwcli.DeleteMessage(ctx, &messagemwpb.MessageReq{
		ID: h.ID,
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
