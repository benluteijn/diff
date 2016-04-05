// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fd_test

import (
	"fmt"
	"math"

	"github.com/gonum/diff/fd"
	"github.com/gonum/matrix/mat64"
)

func ExampleDerivative() {
	f := func(x float64) float64 {
		return math.Sin(x)
	}
	// Compute the first derivative of f at 0 using the default settings.
	fmt.Println("f'(0) ≈", fd.Derivative(f, 0, nil))
	// Compute the first derivative of f at 0 using the forward approximation
	// with a custom step size.
	df := fd.Derivative(f, 0, &fd.Settings{
		Formula: fd.Forward,
		Step:    1e-3,
	})
	fmt.Println("f'(0) ≈", df)

	f = func(x float64) float64 {
		return math.Pow(math.Cos(x), 3)
	}
	// Compute the second derivative of f at 0 using
	// the centered approximation, concurrent evaluation,
	// and a known function value at x.
	df = fd.Derivative(f, 0, &fd.Settings{
		Formula:     fd.Central2nd,
		Concurrent:  true,
		OriginKnown: true,
		OriginValue: f(0),
	})
	fmt.Println("f''(0) ≈", df)

	// Output:
	// f'(0) ≈ 1
	// f'(0) ≈ 0.9999998333333416
	// f''(0) ≈ -2.999999981767587
}

func ExampleJacobian() {
	f := func(y, x []float64) {
		y[0] = x[0] + 1
		y[1] = 5 * x[2]
		y[2] = 4*x[1]*x[1] - 2*x[2]
		y[3] = x[2] * math.Sin(x[0])
	}
	x := []float64{1, 2, 3}
	jac := fd.Jacobian(nil, f, 4, x, &fd.JacobianSettings{
		Formula:    fd.Central,
		Concurrent: true,
	})
	fmt.Printf("J ≈ %v\n", mat64.Formatted(jac, mat64.Prefix("    ")))

	// Output:
	// J ≈ ⎡ 0.9999999999917482                    0                    0⎤
	//     ⎢                  0                    0    4.999999999810711⎥
	//     ⎢                  0   15.999999999719941  -1.9999999999834963⎥
	//     ⎣ 1.6209069175765478                    0   0.8414709847803792⎦
}
