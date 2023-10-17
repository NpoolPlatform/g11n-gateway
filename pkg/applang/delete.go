package applang

import (
	"context"

	applangmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"
	applangmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
)

func (h *Handler) DeleteLang(ctx context.Context) (*npool.Lang, error) {
	exist, err := applangmwcli.ExistAppLangConds(ctx, &applangmwpb.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: *h.ID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: *h.AppID},
	})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, nil
	}

	info, err := applangmwcli.DeleteLang(ctx, &applangmwpb.LangReq{ID: h.ID})
	if err != nil {
		return nil, err
	}

	outs, err := expand(ctx, []*applangmwpb.Lang{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}
