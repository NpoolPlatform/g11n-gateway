package message

import (
	"context"
	"fmt"

	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func (h *Handler) UpdateMessage(ctx context.Context) (*npool.Message, error) {
	exist, err := messagemwcli.ExistMessageConds(ctx, &messagemwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid message")
	}

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
