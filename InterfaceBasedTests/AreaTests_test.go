package interfacebasedtests

import "testing"

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{&Rectangle{12, 6}, 72.0},
		{&Square{5}, 25.0},
		{&Triangle{5, 6}, 15.0},
	}

	for _, testexample := range areaTest {
		got := testexample.shape.Area()
		if got != testexample.want {
			t.Errorf("got %g want %g", got, testexample.want)
		}
	}
}
