package lang

import (
	"context"

	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"
	langmwpb "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
)

func (h *Handler) CreateLang(ctx context.Context) (*langmwpb.Lang, error) {
	return langmwcli.CreateLang(ctx, &langmwpb.LangReq{
		EntID: h.EntID,
		Lang:  h.Lang,
		Logo:  h.Logo,
		Name:  h.Name,
		Short: h.Short,
	})
}

func (h *Handler) CreateLangs(ctx context.Context) ([]*langmwpb.Lang, error) {
	return langmwcli.CreateLangs(ctx, h.Reqs)
}
