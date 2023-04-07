package generator

import (
	"fmt"
	"math"
)

var n float64 = 20.0

const delta float64 = 0.5

var h float64 = 2 * math.Pi / n

var N float64 = (1 * delta) / (math.Sqrt(2) * h)

func deltaFuncUp(x float64) float64 {
	return x*x - math.Pi*math.Pi + N
}

func deltaFuncDown(x float64) float64 {
	return x*x - math.Pi*math.Pi - N
}

func trapezoid_ruleDelta(left float64, right float64, number_of_steps int, number_of_terms int, point_of_research float64, f func(float64) float64, ch chan float64) {
	sum := float64(0)
	h := (right - left) / float64(number_of_steps)
	for i := 1; i < number_of_steps; i++ {
		sum += integration_core(number_of_terms, left+float64(i)*h, point_of_research) * f(left+float64(i)*h)
	}
	sum += integration_core(number_of_terms, left, point_of_research)*f(left) + integration_core(number_of_terms, right, point_of_research)*f(right)/2.0
	sum *= h

	ch <- sum
}

func simpsonRuleDelta(left float64, right float64, number_of_steps int, number_of_terms int, point_of_research float64, f func(float64) float64, ch chan<- float64) {

	sum := float64(0)
	h := (right - left) / (2 * float64(number_of_steps))

	for i := 1; i <= 2*number_of_steps; i++ {
		if i%2 == 0 {
			sum += 2 * integration_core(number_of_terms, left+float64(i)*h, point_of_research) * f(left+float64(i)*h)

		} else {
			sum += 4 * integration_core(number_of_terms, left+float64(i)*h, point_of_research) * f(left+float64(i)*h)

		}
	}
	sum = sum + integration_core(number_of_terms, left, point_of_research)*f(left) + integration_core(number_of_terms, right, point_of_research)*f(right)
	sum = sum * (h / 3)

	ch <- sum
}

func RungeSimpUp(point float64, number_of_terms int) float64 {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)
	go simpsonRuleDelta(-math.Pi, math.Pi, 5, number_of_terms, point, deltaFuncUp, ch)
	step = <-ch
	go simpsonRuleDelta(-math.Pi, math.Pi, 10, number_of_terms, point, deltaFuncUp, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go simpsonRuleDelta(-math.Pi, math.Pi, n, number_of_terms, point, deltaFuncUp, ch)
		step = <-ch

		go simpsonRuleDelta(-math.Pi, math.Pi, 2*n, number_of_terms, point, deltaFuncUp, ch)
		half_step = <-ch

	}
	fmt.Println(step)
	simpResult := step

	return simpResult
}

func RungeTrapUp(point float64, number_of_terms int) float64 {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)

	go trapezoid_ruleDelta(-math.Pi, math.Pi, 2, number_of_terms, point, deltaFuncUp, ch)
	step = <-ch
	fmt.Println(step)
	go trapezoid_ruleDelta(-math.Pi, math.Pi, 4, number_of_terms, point, deltaFuncUp, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go trapezoid_ruleDelta(-math.Pi, math.Pi, n, number_of_terms, point, deltaFuncUp, ch)
		step = <-ch

		go trapezoid_ruleDelta(-math.Pi, math.Pi, 2*n, number_of_terms, point, deltaFuncUp, ch)
		half_step = <-ch

		fmt.Println("step:", step, half_step, "parameter:", n, number_of_terms, point)
	}

	trapResult := step
	return trapResult
}

func RungeSimpDown(point float64, number_of_terms int) float64 {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)
	go simpsonRuleDelta(-math.Pi, math.Pi, 5, number_of_terms, point, deltaFuncDown, ch)
	step = <-ch
	go simpsonRuleDelta(-math.Pi, math.Pi, 10, number_of_terms, point, deltaFuncDown, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go simpsonRuleDelta(-math.Pi, math.Pi, n, number_of_terms, point, deltaFuncDown, ch)
		step = <-ch

		go simpsonRuleDelta(-math.Pi, math.Pi, 2*n, number_of_terms, point, deltaFuncDown, ch)
		half_step = <-ch

	}
	fmt.Println(step)
	simpResult := step

	return simpResult
}

func RungeTrapDown(point float64, number_of_terms int) float64 {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)

	go trapezoid_ruleDelta(-math.Pi, math.Pi, 2, number_of_terms, point, deltaFuncDown, ch)
	step = <-ch
	fmt.Println(step)
	go trapezoid_ruleDelta(-math.Pi, math.Pi, 4, number_of_terms, point, deltaFuncDown, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go trapezoid_ruleDelta(-math.Pi, math.Pi, n, number_of_terms, point, deltaFuncDown, ch)
		step = <-ch

		go trapezoid_ruleDelta(-math.Pi, math.Pi, 2*n, number_of_terms, point, deltaFuncDown, ch)
		half_step = <-ch

		fmt.Println("step:", step, half_step, "parameter:", n, number_of_terms, point)
	}

	trapResult := step
	return trapResult
}
