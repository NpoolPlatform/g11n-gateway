package lang

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	langmw "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
	"github.com/google/uuid"
)

type Handler struct {
	ID     *uint32
	EntID  *string
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

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithLang(lang *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lang == nil {
			if must {
				return fmt.Errorf("invalid lang")
			}
			return nil
		}
		if *lang == "" {
			return fmt.Errorf("invalid lang")
		}
		h.Lang = lang
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid langname")
		}
		h.Name = name
		return nil
	}
}

func WithLogo(logo *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if logo == nil {
			if must {
				return fmt.Errorf("invalid logo")
			}
			return nil
		}
		if *logo == "" {
			return fmt.Errorf("invalid logo")
		}
		h.Logo = logo
		return nil
	}
}

func WithShort(short *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if short == nil {
			if must {
				return fmt.Errorf("invalid short")
			}
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
		if len(reqs) == 0 {
			return fmt.Errorf("infos is empty")
		}
		for _, req := range reqs {
			if req.EntID != nil {
				_, err := uuid.Parse(*req.EntID)
				if err != nil {
					return err
				}
			}
			if req.Lang == nil || *req.Lang == "" {
				return fmt.Errorf("invalid lang")
			}
		}
		h.Reqs = reqs
		return nil
	}
}
