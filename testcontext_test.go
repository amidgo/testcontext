package testcontext

import "testing"

func Test_Context_Lookup_Assert(t *testing.T) {
	ctx := Context(t)

	ctxT, ok := Lookup(ctx)

	if ctxT != t {
		t.Fatal("Lookup testing, invalid *testing.T")
	}

	if !ok {
		t.Fatal("Lookup testing, invalid bool")
	}

	inc := 0

	Test(ctx, func(t *testing.T) { inc++ })

	if inc != 1 {
		t.Fatalf("Assert testing, invalid inc value, expected 1, actual %d", inc)
	}
}
