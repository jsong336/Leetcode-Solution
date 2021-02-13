package main

import "fmt"

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	var two_index [] int
	for i, n_i := range nums {
		for j, n_j := range nums{
			if (n_i + n_j == target) && (i != j){
				two_index = [] int {i, j}
				return two_index
			}
		}
	}
	two_index = [] int {}
	return two_index
}

func main() {
	nums := []int{3,2,4}
	target := 6
	answer := twoSum(nums, target)
	fmt.Println(answer)
}