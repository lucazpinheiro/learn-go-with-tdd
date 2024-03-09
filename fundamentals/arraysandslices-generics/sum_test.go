package arraysandslicesgenerics

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestSum(t *testing.T) {
	t.Run("collection of n numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %v want %v given, %v", got, want, numbers)
		}
	})

}

func TestSumAllTail(t *testing.T) {

	/*
		We could've created a new function checkSums like we normally do, but in this case,
		we're showing a new technique, assigning a function to a variable. It might look
		strange but, it's no different to assigning a variable to a string, or an int,
		functions in effect are values too.

		It's not shown here, but this technique can be useful when you want to bind a function
		to other local variables in "scope" (e.g between some {}). It also allows you to reduce
		the surface area of your API.

		By defining this function inside the test, it cannot be used by other functions in this
		package. Hiding variables and functions that don't need to be exported is an important
		design consideration.

		A handy side-effect of this is this adds a little type-safety to our code. If a
		developer mistakenly adds a new test with checkSums(t, got, "dave") the compiler will
		stop them in their tracks.
	*/
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		// reflect.DeepEqual is a builtin function to compare arrays, slices and other things
		// it is necessary since the equality operators (==, !=) don't work for slices and arrays
		// IMPORTANT: it is not type safe, the compiler will compile even if a string and slice are
		// passed as arguments
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	}
	t.Run("make the sums of  some slices", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{2, 4}
		s3 := []int{3, 4, 10}

		got := SumAllTails(s1, s2, s3)
		want := []int{5, 4, 14}

		// reflect.DeepEqual is a builtin function to compare arrays, slices and other things
		// it is necessary since the equality operators (==, !=) don't work for slices and arrays
		// IMPORTANT: it is not type safe, the compiler will compile even if a string and slice are
		// passed as arguments
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		s1 := []int{}
		s2 := []int{2, 4}
		s3 := []int{3, 4, 10}

		got := SumAllTails(s1, s2, s3)
		want := []int{0, 4, 14}

		checkSums(t, got, want)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
