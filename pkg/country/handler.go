package country

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	countrymw "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uint32
	EntID   *string
	Country *string
	Flag    *string
	Code    *string
	Short   *string
	Reqs    []*countrymw.CountryReq
	Offset  int32
	Limit   int32
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

func WithCountry(country *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if country == nil {
			if must {
				return fmt.Errorf("invalid country")
			}
			return nil
		}
		if *country == "" {
			return fmt.Errorf("invalid country")
		}
		h.Country = country
		return nil
	}
}

func WithFlag(flag *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if flag == nil {
			if must {
				return fmt.Errorf("invalid flag")
			}
			return nil
		}
		if *flag == "" {
			return fmt.Errorf("invalid flag")
		}
		h.Flag = flag
		return nil
	}
}

func WithCode(code *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if code == nil {
			if must {
				return fmt.Errorf("invalid code")
			}
			return nil
		}
		if *code == "" {
			return fmt.Errorf("invalid code")
		}
		h.Code = code
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

func WithReqs(reqs []*countrymw.CountryReq) func(context.Context, *Handler) error {
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
			if req.Country == nil || *req.Country == "" {
				return fmt.Errorf("invalid country")
			}
		}
		h.Reqs = reqs
		return nil
	}
}
