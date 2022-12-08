package lang

import (
	"context"

	langmgrcli "github.com/NpoolPlatform/g11n-manager/pkg/client/lang"
	langmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"
)

func GetLangs(ctx context.Context, in *langmgrpb.Conds, offset, limit int32) ([]*langmgrpb.Lang, uint32, error) {
	return langmgrcli.GetLangs(ctx, in, offset, limit)
}
