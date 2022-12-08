package lang

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"
)

func GetLangs(ctx context.Context, in *langmgrpb.Conds) ([]*npool.Lang, uint32, error) {
	return nil, 0, nil
}
