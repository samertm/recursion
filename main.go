// recursion practice

package main

import (
	"fmt"
)

// TODO time complexity
func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n - 1) + fib(n - 2)
}

// TODO time complexity
func max(nums []int) int {
	var max_nums func([]int, int) int
	max_nums = func(nums []int, largest int) int {
		if len(nums) == 0 {
			return largest
		}
		next := nums[0]
		if next > largest {
			return max_nums(nums[1:], next)
		} else {
			return max_nums(nums[1:], largest)
		}		
	}
	return max_nums(nums, 0)
}

// TODO time complexity
func sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + sum(nums[1:])
}

// TODO time complexity
func lastIndexOf(n int, nums []int) int {
	var helper func(int,int,[]int) int 
	helper = func(c, n int, nums []int) int {
		if len(nums) == 0 {
			return -1
		}
		var indexVal int
		if nums[0] == n {
			indexVal = c
		} else {
			indexVal = -1
		}		
		bubbleIndex := helper(c+1, n, nums[1:])
		if bubbleIndex > indexVal {
			return bubbleIndex
		} else {
			return indexVal
		}		
	}
	return helper(0, n, nums)
}

func main() {
	//fmt.Println(fib(10))
	//fmt.Println(max([]int{4,6,7,8,33,54,42,1}))
	//fmt.Println(sum([]int{4,6,7,8,33,54,42,1}))
	fmt.Println(lastIndexOf(4, []int{4,6,7,8,33,54,8,42,1}))
	fmt.Println(lastIndexOf(8, []int{4,6,7,8,33,54,8,42,1}))
}
