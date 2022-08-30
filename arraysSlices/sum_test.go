package summation

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size of numbers", func(t *testing.T) {
		summation := []int{1, 2, 3}

		got := Sum(summation)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given", got, want)
		}
	})

}
