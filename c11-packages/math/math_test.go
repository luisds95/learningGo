package math

import "testing"

type testData struct {
	values []float64
	mean float64
}

var tests = []testData{
	{ []float64{5, 10, 15, 20}, 12.5 },
	{ []float64{1, 1, 1, 1, 1, 1}, 1},
	{ []float64{-10, 6}, -2 },
	{ []float64{}, 0 }
}

func TestAverage(t *testing.T) {
	for _, data := range tests {
		avg := Average(data.values)
		if avg != data.avg {
			t.Error(
				"For", data.values,
				"expected", data.mean,
				"but got", avg
			)
		}
	}
}