package test

import (
	"fmt"
	"math"
	"testing"
	"gostudy/program"
)

func TestEval(t *testing.T)  {
	tests := []struct{
		expr string
		env program.Env
		want string
	}{
		{"sqrt(A / pi)", program.Env{"A": 87617, "pi": math.Pi}, "167"},
	}

	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}

		expr, err := program.Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Println(got)


	}
}
