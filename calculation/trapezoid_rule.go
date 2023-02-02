package calculation

func trapezoid_rule(left float32, right float32, number_of_steps int) float32 {
	sum := float32(0)
	h := (right - left) / float32(number_of_steps)
	for i := 1; i < number_of_steps; i++ {
		sum += base_function(left + float32(i)*h)
	}
	sum += (base_function(left) + base_function(right)) / 2.0
	sum *= h
	return sum
}
