package testcontext

import (
	"context"
	"testing"
)

type testingKey struct{}

func Context(t *testing.T) context.Context {
	return context.WithValue(t.Context(), testingKey{}, t)
}

func Lookup(ctx context.Context) (*testing.T, bool) {
	t, ok := ctx.Value(testingKey{}).(*testing.T)

	return t, ok
}

func Assert(ctx context.Context, asrt func(t *testing.T)) {
	t, ok := Lookup(ctx)
	if ok {
		asrt(t)
	}
}
