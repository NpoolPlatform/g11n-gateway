package message

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	constant "github.com/NpoolPlatform/g11n-gateway/pkg/const"
	langmwcli "github.com/NpoolPlatform/g11n-middleware/pkg/client/lang"
	messagemw "github.com/NpoolPlatform/message/npool/g11n/mw/v1/message"
	"github.com/google/uuid"
)

type Handler struct {
	ID        *string
	AppID     *string
	LangID    *string
	MessageID *string
	Message   *string
	GetIndex  *uint32
	Disabled  *bool
	Reqs      []*messagemw.MessageReq
	Offset    int32
	Limit     int32
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

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_app, err := appmwcli.GetApp(ctx, *id)
		if err != nil {
			return err
		}
		if _app == nil {
			return fmt.Errorf("invalid app")
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
		_app, err := langmwcli.GetLang(ctx, *id)
		if err != nil {
			return err
		}
		if _app == nil {
			return fmt.Errorf("invalid langid")
		}
		h.LangID = id
		return nil
	}
}

func WithMessageID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MessageID = id
		return nil
	}
}

func WithMessage(message *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Message = message
		return nil
	}
}

func WithGetIndex(getindex *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GetIndex = getindex
		return nil
	}
}

func WithDisabled(disabled *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Disabled = disabled
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

func WithReqs(reqs []*messagemw.MessageReq) func(context.Context, *Handler) error {
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
			if req.AppID != nil {
				_, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid appid")
			}
			if req.LangID != nil {
				_, err := uuid.Parse(*req.LangID)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid langid")
			}
			if req.MessageID == nil {
				return fmt.Errorf("invalid messageid")
			}
		}
		h.Reqs = reqs
		return nil
	}
}

func WithAppReqs(reqs []*messagemw.MessageReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(reqs) == 0 {
			return fmt.Errorf("infos is empty")
		}
		if h.AppID == nil || *h.AppID == "" {
			return fmt.Errorf("invalid targetappid")
		}
		if h.LangID == nil || *h.LangID == "" {
			return fmt.Errorf("invalid targetlangid")
		}
		_reqs := []*messagemw.MessageReq{}
		for _, req := range reqs {
			_req := req
			if req.ID != nil {
				_, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
			}
			if req.MessageID == nil {
				return fmt.Errorf("invalid messageid")
			}
			_req.AppID = h.AppID
			_req.LangID = h.LangID
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
