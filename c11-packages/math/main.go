package math

func Average(numbers []float64) float64 {
	total := float64(0)
	for _, value := range numbers {
		total += value
	}
	return total / float64(len(numbers))
}