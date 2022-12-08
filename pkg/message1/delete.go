package message

import (
	"context"

	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
)

func DeleteMessage(ctx context.Context, id string) (*npool.Message, error) {
	info, err := messagemwcli.DeleteMessage(ctx, id)
	if err != nil {
		return nil, err
	}

	outs, err := Expand(ctx, []*messagemwpb.Message{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}
