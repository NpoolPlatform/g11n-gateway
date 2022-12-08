package message

import (
	"context"

	messagemwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/message"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/message"
	messagemgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
)

func GetMessages(ctx context.Context, in *messagemgrpb.Conds, offset, limit int32) ([]*npool.Message, uint32, error) {
	infos, total, err := messagemwcli.GetMessages(ctx, in, offset, limit)
	if err != nil {
		return nil, 0, nil
	}

	outs, err := Expand(ctx, infos)
	if err != nil {
		return nil, 0, err
	}

	return outs, total, nil
}
