package generator

import (
	"fmt"
	"math"
)

const eps = 0.01

func Base_function(x float64) float64 {
	return x*x - math.Pi*math.Pi
}

func integration_core(number_of_terms int, point float64, static_point float64) float64 {
	sum := 0.0
	for i := 1; i <= number_of_terms; i++ {
		sum += math.Cos(float64(i) * (point - static_point))
	}

	return (1.0 / math.Pi) * (0.5 + sum)
}

func trapezoid_rule(left float64, right float64, number_of_steps int, number_of_terms int, point_of_research float64, ch chan float64) {
	sum := float64(0)
	h := (right - left) / float64(number_of_steps)
	for i := 1; i < number_of_steps; i++ {
		sum += integration_core(number_of_terms, left+float64(i)*h, point_of_research) * Base_function(left+float64(i)*h)
	}
	sum += integration_core(number_of_terms, left, point_of_research)*Base_function(left) + integration_core(number_of_terms, right, point_of_research)*Base_function(right)/2.0
	sum *= h

	ch <- sum
}

func simpsonRule(left float64, right float64, number_of_steps int, number_of_terms int, point_of_research float64, ch chan<- float64) {

	sum := float64(0)
	h := (right - left) / (2 * float64(number_of_steps))

	for i := 1; i <= 2*number_of_steps; i++ {
		if i%2 == 0 {
			sum += 2 * integration_core(number_of_terms, left+float64(i)*h, point_of_research) * Base_function(left+float64(i)*h)

		} else {
			sum += 4 * integration_core(number_of_terms, left+float64(i)*h, point_of_research) * Base_function(left+float64(i)*h)

		}
	}
	sum = sum + integration_core(number_of_terms, left, point_of_research)*Base_function(left) + integration_core(number_of_terms, right, point_of_research)*Base_function(right)
	sum = sum * (h / 3)

	ch <- sum
}

func RungeSimp(point float64, number_of_terms int) float64 {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)
	go simpsonRule(-math.Pi, math.Pi, 5, number_of_terms, point, ch)
	step = <-ch
	go simpsonRule(-math.Pi, math.Pi, 10, number_of_terms, point, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go simpsonRule(-math.Pi, math.Pi, n, number_of_terms, point, ch)
		step = <-ch

		go simpsonRule(-math.Pi, math.Pi, 2*n, number_of_terms, point, ch)
		half_step = <-ch

	}
	fmt.Println(step)
	simpResult := step

	return simpResult
}

func RungeTrap(point float64, number_of_terms int) float64 {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)

	go trapezoid_rule(-math.Pi, math.Pi, 2, number_of_terms, point, ch)
	step = <-ch
	fmt.Println(step)
	go trapezoid_rule(-math.Pi, math.Pi, 4, number_of_terms, point, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go trapezoid_rule(-math.Pi, math.Pi, n, number_of_terms, point, ch)
		step = <-ch

		go trapezoid_rule(-math.Pi, math.Pi, 2*n, number_of_terms, point, ch)
		half_step = <-ch

		fmt.Println("step:", step, half_step, "parameter:", n, number_of_terms, point)
	}

	trapResult := step
	return trapResult
}
