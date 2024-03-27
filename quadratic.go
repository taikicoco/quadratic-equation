package quadratic

import (
	"math/cmplx"
)

func Solve(a, b, c float64) (x1, x2 complex128) {
	D := b*b - 4*a*c
	sqrtD := cmplx.Sqrt(complex(D, 0))

	x1 = (-complex(b, 0) + sqrtD) / complex(2*a, 0)
	x2 = (-complex(b, 0) - sqrtD) / complex(2*a, 0)

	return x1, x2
}
