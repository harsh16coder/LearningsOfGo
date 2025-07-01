package sumall

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSum := func(got, expected []int, t testing.TB) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got this %v expected this %v", got, expected)
		}
	}
	t.Run("Make sum of all elements", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{5, 4, 9})
		expected := []int{6, 18}
		checkSum(got, expected, t)
	})
	t.Run("pass empty slice", func(t *testing.T) {
		got := SumAll([]int{}, []int{5, 4, 9})
		expected := []int{0, 18}
		checkSum(got, expected, t)
	})
}
