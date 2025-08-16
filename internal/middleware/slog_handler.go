package middleware

import (
	"context"
	"log/slog"
)

type Handler struct{ h slog.Handler }

func New(h slog.Handler) *Handler { return &Handler{h: h} }

func (c *Handler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return c.h.Enabled(ctx, lvl)
}

func (c *Handler) Handle(ctx context.Context, r slog.Record) error {
	if attrs := Get(ctx); len(attrs) > 0 {
		r.AddAttrs(attrs...)
	}
	return c.h.Handle(ctx, r)
}

func (c *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: c.h.WithAttrs(attrs)}
}

func (c *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: c.h.WithGroup(name)}
}

type key struct{}

func Add(ctx context.Context, attrs ...slog.Attr) context.Context {
	if len(attrs) == 0 {
		return ctx
	}
	existing, _ := ctx.Value(key{}).([]slog.Attr)
	out := make([]slog.Attr, 0, len(existing)+len(attrs))
	out = append(out, existing...)
	out = append(out, attrs...)
	return context.WithValue(ctx, key{}, out)
}

// Get extracts all attrs previously stored with Add.
func Get(ctx context.Context) []slog.Attr {
	attrs, _ := ctx.Value(key{}).([]slog.Attr)
	return attrs
}
