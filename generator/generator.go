package generator

import (
	"math"
)

type pointFromMethods struct {
	simpsonPoint   float64
	trapezoidPoint float64
}

const start_pos = -math.Pi
const eps = 0.001

func Base_function(x float64) float64 {
	return x*x - math.Pi*math.Pi
}

func integration_core(number_of_terms int, point float64, static_point float64) float64 {
	sum := 0.0
	for i := 0; i < number_of_terms; i++ {
		sum += math.Cos(point - static_point)
	}

	return (1.0 / math.Pi) * (0.5 + sum)
}

func trapezoid_rule(left float64, right float64, number_of_steps int, number_of_terms int, point_of_research float64, ch chan float64) {
	sum := float64(0)
	h := (right - left) / float64(number_of_steps)
	for i := 1; i < number_of_steps; i++ {
		sum += Base_function(left + float64(i)*h)
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

func runge_rule(point float64, number_of_steps int) pointFromMethods {
	step := float64(0)
	half_step := float64(0)
	ch := make(chan float64, 2)
	go simpsonRule(-math.Pi, math.Pi, number_of_steps, 5, 0, ch)
	step = <-ch
	go simpsonRule(-math.Pi, math.Pi, number_of_steps, 5, 0, ch)
	half_step = <-ch

	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go simpsonRule(-math.Pi, math.Pi, number_of_steps, n, point, ch)
		step = <-ch
		go simpsonRule(-math.Pi, math.Pi, number_of_steps, 2*n, point, ch)
		half_step = <-ch
	}
	simpResult := step
	step = float64(0)
	half_step = float64(0)

	go trapezoid_rule(-math.Pi, math.Pi, number_of_steps, 5, 0, ch)
	step = <-ch
	go trapezoid_rule(-math.Pi, math.Pi, number_of_steps, 5, 0, ch)
	half_step = <-ch
	for n := 5; math.Abs(step-half_step) >= eps; n++ {
		go trapezoid_rule(-math.Pi, math.Pi, number_of_steps, n, point, ch)
		step = <-ch
		go trapezoid_rule(-math.Pi, math.Pi, number_of_steps, 2*n, point, ch)
		half_step = <-ch
	}
	trapResult := step
	var points pointFromMethods = pointFromMethods{simpsonPoint: simpResult, trapezoidPoint: trapResult}

	return points
}
