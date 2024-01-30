package main

import utl "autilities"

func addAll(nums ...int) int {
	utl.PLine("Numbers: ", nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	utl.PLine("Total: ", total)

	return total
}
