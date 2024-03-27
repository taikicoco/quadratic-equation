package quadratic

import (
	"math/cmplx"
	"testing"
)

func TestQuadratic(t *testing.T) {
	tests := []struct {
		a, b, c        float64
		wantX1, wantX2 complex128
	}{
		{1, 3, 2, -1, -2},
		{1, 2, 1, -1, -1},
		{1, 0, 1, cmplx.Sqrt(-1), -cmplx.Sqrt(-1)},
	}

	for _, tt := range tests {
		gotX1, gotX2 := Solve(tt.a, tt.b, tt.c)
		if gotX1 != tt.wantX1 || gotX2 != tt.wantX2 {
			t.Errorf("Quadratic(%v, %v, %v) = (%v, %v), want (%v, %v)",
				tt.a, tt.b, tt.c, gotX1, gotX2, tt.wantX1, tt.wantX2)
		}
	}
}
