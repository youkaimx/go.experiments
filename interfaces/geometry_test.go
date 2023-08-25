package interfaces

import "testing"

const (
	PERIMETER = 0
	AREA      = 1
)

func Test_PrintPerimeterAndArea(t *testing.T) {
	test_cases := []struct {
		name     string
		shape    TwoDShape
		expected []float64
	}{
		{
			"Rectangle",
			Rectangle{[]float64{8, 6}},
			[]float64{28, 48},
		},
		{
			"Square",
			Square{10},
			[]float64{40, 100},
		},
	}

	for _, test_case := range test_cases {
		t.Run(test_case.name, func(t *testing.T) {
			perimeter, area := getPerimeterAndArea(test_case.shape)
			if perimeter != test_case.expected[PERIMETER] {
				t.Errorf("Perimeter expected: %f, got: %f", test_case.expected[PERIMETER], perimeter)
			}
			if area != test_case.expected[AREA] {
				t.Errorf("Area expected: %f, got: %f", test_case.expected[AREA], area)
			}
		})
	}
}
