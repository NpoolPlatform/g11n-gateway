package applang

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"

	"github.com/google/uuid"
)

type Handler struct {
	ID     *string
	AppID  string
	LangID *string
	Main   *bool
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

func WithAppID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		exist, err := appmwcli.ExistApp(ctx, id)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("app not exist")
		}
		h.AppID = id
		return nil
	}
}

func WithLangID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_lang, err := langmwcli.GetLang(ctx, *id)
		if err != nil {
			return err
		}
		if _lang == nil {
			return fmt.Errorf("invalid lang")
		}
		h.LangID = id
		return nil
	}
}

func WithMain(main *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if main == nil {
			return nil
		}
		h.Main = main
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
