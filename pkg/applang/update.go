package applang

import (
	"context"

	applangmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"
	applangmgrpb "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
	applangmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
)

func UpdateLang(ctx context.Context, in *applangmgrpb.LangReq) (*npool.Lang, error) {
	info, err := applangmwcli.UpdateLang(ctx, in)
	if err != nil {
		return nil, err
	}

	outs, err := expand(ctx, []*applangmwpb.Lang{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}
