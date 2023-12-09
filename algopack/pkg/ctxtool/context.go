package ctxtool

import (
	"context"

	"go.uber.org/zap"
)

const (
	ctxStoreKey ctxMarker = iota
)

type ctxMarker uint8

type ctxStore struct {
	logger *zap.Logger
}

func newCtxStore() *ctxStore {
	return &ctxStore{}
}

func withCtxStore(ctx context.Context) (context.Context, *ctxStore) {
	store, ok := getCtxStore(ctx)
	if !ok {
		ctx = context.WithValue(ctx, ctxStoreKey, store)
	}
	return ctx, store
}

func getCtxStore(ctx context.Context) (*ctxStore, bool) {
	v := ctx.Value(ctxStoreKey)
	if v == nil {
		return newCtxStore(), false
	}

	return v.(*ctxStore), true
}

func Logger(ctx context.Context) *zap.Logger {
	store, _ := getCtxStore(ctx)
	if store.logger != nil {
		return store.logger
	}

	return zap.L()
}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	resCtx, store := withCtxStore(ctx)
	store.logger = logger

	return resCtx
}
