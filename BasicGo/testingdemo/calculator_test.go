package testingdemo

import (
	"math/rand"
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		name string
		a, b float64
		op   Operation
		want float64
		err  error
	}{
		{"add", 1, 2, Add, 3, nil},
		{"sub", 5, 2, Sub, 3, nil},
		{"mul", 3, 3, Mul, 9, nil},
		{"div", 10, 2, Div, 5, nil},
		{"div-by-zero", 10, 0, Div, 0, ErrDivideByZero},
		{"unknown", 1, 1, Operation("pow"), 0, ErrUnknownOp},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calc(tc.a, tc.b, tc.op)
			if err != tc.err {
				t.Fatalf("expected err %v, got %v", tc.err, err)
			}
			if got != tc.want {
				t.Fatalf("expected result %f, got %f", tc.want, got)
			}
		})
	}
}

func BenchmarkCalc(b *testing.B) {
	ops := []Operation{Add, Sub, Mul, Div}
	for i := 0; i < b.N; i++ {
		op := ops[rand.Intn(len(ops))]
		_, _ = Calc(10, 2, op)
	}
}
