package lang

import (
	"context"

	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"
	langmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

func (h *Handler) GetLangs(ctx context.Context) ([]*langmwpb.Lang, uint32, error) {
	return langmwcli.GetLangs(ctx, &langmwpb.Conds{}, h.Offset, h.Limit)
}
