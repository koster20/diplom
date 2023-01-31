package calculation

import (
	"math"
)

const start_pos = -math.Pi

func base_function(x float32) float32 {
	return x*x - math.Pi*math.Pi
}
