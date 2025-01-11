package main

import (
	"fmt"
	"strconv"
)

func generateNumberSet(pattern string) string {
	n := len(pattern) + 1
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = 0
	}
	minSum := -1
	var bestResult string

	for {
		if isValid(nums, pattern) {
			result := numsToString(nums)
			sum := sumOfDigits(result)
			if minSum == -1 || sum < minSum {
				minSum = sum
				bestResult = result
			}
		}
		if !increment(nums) {
			break
		}
	}
	return bestResult
}

func isValid(nums []int, pattern string) bool {
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case 'L':
			if nums[i] <= nums[i+1] {
				return false
			}
		case 'R':
			if nums[i] >= nums[i+1] {
				return false
			}
		case '=':
			if nums[i] != nums[i+1] {
				return false
			}
		}
	}
	return true
}

func increment(nums []int) bool {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < len(nums)-1 {
			nums[i]++
			return true
		}
		nums[i] = 0
	}
	return false
}

func numsToString(nums []int) string {
	var result string
	for _, num := range nums {
		result += strconv.Itoa(num)
	}
	return result
}

func sumOfDigits(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		sum += num
	}
	return sum
}

func main() {
	var input string
	fmt.Print("Enter the pattern (e.g., LLRR=): ")
	fmt.Scanln(&input)

	result := generateNumberSet(input)
	fmt.Printf("output: %s\n", result)
}
