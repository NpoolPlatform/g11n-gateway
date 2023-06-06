package lang

import (
	"context"

	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"
	langmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

func (h *Handler) UpdateLang(ctx context.Context) (*langmwpb.Lang, error) {
	return langmwcli.UpdateLang(ctx, &langmwpb.LangReq{
		ID:    h.ID,
		Lang:  h.Lang,
		Logo:  h.Logo,
		Name:  h.Name,
		Short: h.Short,
	})
}
