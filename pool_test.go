package pool

import (
	"context"
	"testing"
)

func TestPool(t *testing.T) {
	var cnt int
	p := Pool[int]{
		New: func(_ context.Context) (int, error) {
			cnt++
			return cnt, nil
		},
	}
	defer p.Close()

	ctx := context.Background()
	x, err := p.Get(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if x != 1 {
		t.Fatalf("got %d, want 1", x)
	}
	p.Put(x)

	x, err = p.Get(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if x != 1 {
		t.Fatalf("got %d, want 1", x)
	}
	p.Put(x)
}
