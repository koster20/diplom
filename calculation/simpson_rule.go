package calculation

func simpson_rule(left float32, right float32, number_of_steps int) float32 {

	sum := float32(0)
	h := (right - left) / (2 * float32(number_of_steps))

	for i := 1; i <= 2*number_of_steps; i++ {
		if i%2 == 0 {
			sum += 2 * base_function(left+float32(i)*h)

		} else {
			sum += 4 * base_function(left+float32(i)*h)

		}
	}
	sum = sum + base_function(left) + base_function(right)
	sum = sum * (h / 3)

	return sum
}
