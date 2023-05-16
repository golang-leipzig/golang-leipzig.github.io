// Write a function in Go to check, whether a sequence of numbers contains two numbers that sum to zero.

package main

import "fmt"

func ContainsSumZero(nums []int) bool {
	seen := make(map[int]bool) // keep track of already encountered nums
	for _, num := range nums {
		complement := 0 - num // try every combination of seen/unseen numbers
		if seen[complement] {
			fmt.Println("Found pair:", num, complement)
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	test_list := []int{34, 78, -90, 65, 23, 12, 90, 34, 89, 23, 100}
	found := ContainsSumZero(test_list)
	if !found {
		fmt.Println("No matching pair found")
	}
}

// In this program, ContainsSumZero takes a slice of ints as input and returns
// true if there exists any pair of elements whose sum is equal to 0. It does so
// by first creating a map called seen, which maps each unencountered integer to
// boolean True. Then, it iterates over each element in the input slice and tries
// adding both positive and negative versions of said number to see if they appear
// earlier in the list (if seen[complement]). If yes, then it prints out those two
// numbers alongside each other and returns early indicating success via return
// true. Otherwise, it updates the mapping for the current number being considered
// (seen[num]=true). Finally, after exhausting the entire input list without
// finding a match, it returns False.
