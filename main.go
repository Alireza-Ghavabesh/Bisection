package main

import "fmt"

var n int = 0

func f(x float64) float64 {
	return (x * x) - (3 * x) - 4
}

func tansif(a float64, b float64, f func(float64) float64) {
	n++
	var x float64 = (b - a) / 2
	fmt.Printf("x_%d = %f\n", n, x)
	if f(a)*f(x) < 0 {
		b = x
	} else if f(x)*f(b) < 0 {
		a = x
	}
	fmt.Printf("a = %f\n", a)
	fmt.Printf("b = %f\n", b)
	fmt.Printf("===========================\n")
	if n >= 21 {
		return
	} else {
		tansif(a, b, f)
	}
}

func main() {
	tansif(-2.0, 1.0, f)
	fmt.Println(n)
}
