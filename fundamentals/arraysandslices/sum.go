package arraysandslices

func Sum(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAllTails(slices ...[]int) []int {
	tailsSlices := []int{}

	for _, s := range slices {
		if len(s) == 0 {
			tailsSlices = append(tailsSlices, 0)
			continue
		}
		tailsSlices = append(tailsSlices, Sum(s[1:]))
	}

	return tailsSlices
}

// function SumALL will be changed to SumAllTails
// func SumAll(slices ...[]int) []int {
// 	// resultingSlice := []int{} 1° option to start an empty slice

// 	// start an empty slice with and predefined capacity
// 	resultingSlice := make([]int, len(slices)) // 1° option to start an empty slice

// 	for _, s := range slices {
// 		resultingSlice = append(resultingSlice, Sum(s))
// 	}

// 	return resultingSlice
// }
