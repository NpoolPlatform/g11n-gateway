package lang

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	langmw "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
	"github.com/google/uuid"
)

type Handler struct {
	ID     *string
	Lang   *string
	Name   *string
	Logo   *string
	Short  *string
	Reqs   []*langmw.LangReq
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.ID = id
		return nil
	}
}

func WithLang(lang *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lang == nil {
			return nil
		}
		_lang, err := langmwcli.GetLangOnly(ctx, &langmw.Conds{
			Lang: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: *lang,
			},
		})
		if err != nil {
			return err
		}
		if _lang == nil {
			return fmt.Errorf("invalid lang")
		}
		h.Lang = lang
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid langname")
		}
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if logo == nil {
			return nil
		}
		if *logo == "" {
			return fmt.Errorf("invalid logo")
		}
		h.Logo = logo
		return nil
	}
}

func WithShort(short *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if short == nil {
			return nil
		}
		if *short == "" {
			return fmt.Errorf("invalid short")
		}
		h.Short = short
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}

func WithReqs(reqs []*langmw.LangReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			if req.ID != nil {
				_, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
			}
		}
		h.Reqs = reqs
		return nil
	}
}
