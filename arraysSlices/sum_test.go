package summation

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size of numbers", func(t *testing.T) {
		summation := []int{1, 2, 3}

		got := Sum(summation)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

	t.Run("summation of all", func(t *testing.T) {

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
