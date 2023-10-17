package applang

import (
	"context"

	applangmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/applang"
	npool "github.com/NpoolPlatform/message/npool/g11n/gw/v1/applang"
	applangmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/applang"
)

func (h *Handler) CreateLang(ctx context.Context) (*npool.Lang, error) {
	info, err := applangmwcli.CreateLang(ctx, &applangmwpb.LangReq{
		AppID:  h.AppID,
		LangID: h.LangID,
		Main:   h.Main,
	})
	if err != nil {
		return nil, err
	}

	outs, err := expand(ctx, []*applangmwpb.Lang{info})
	if err != nil {
		return nil, err
	}

	return outs[0], nil
}
