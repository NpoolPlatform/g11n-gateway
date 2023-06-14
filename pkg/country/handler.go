package country

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	countrymw "github.com/NpoolPlatform/message/npool/g11n/mw/v1/country"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *string
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

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = id
		return nil
	}
}

func WithCountry(country *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if country == nil {
			return nil
		}
		if *country == "" {
			return fmt.Errorf("invalid country")
		}
		h.Country = country
		return nil
	}
}

func WithFlag(flag *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if flag == nil {
			return nil
		}
		if *flag == "" {
			return fmt.Errorf("invalid flag")
		}
		h.Flag = flag
		return nil
	}
}

func WithCode(code *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if code == nil {
			return nil
		}
		if *code == "" {
			return fmt.Errorf("invalid code")
		}
		h.Code = code
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

func WithReqs(reqs []*countrymw.CountryReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(reqs) == 0 {
			return fmt.Errorf("infos is empty")
		}
		for _, req := range reqs {
			if req.ID != nil {
				_, err := uuid.Parse(*req.ID)
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
