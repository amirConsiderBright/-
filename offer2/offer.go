package main

import "fmt"

//    找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
// 找出两个元素之和等于8的下标分别是（0，4）和（1，2）

// 求元素和，是给定的值
func myTest(a [5]int, target int) {
	ints := make(map[int]*int)
	for i := 0; i < len(a); i++ {
		i2, ok := ints[target-a[i]]
		if ok {
			fmt.Println("{", i, ",", *i2, "}")
		}
		i3 := i
		ints[a[i]] = &i3
	}
}
func twoSum(nums []int, target int) []int {
	ints := make(map[int]*int)
	for i := 0; i < len(nums); i++ {
		i2, ok := ints[target-nums[i]]
		if ok {
			return []int{i, *i2}
		}
		i3 := i
		ints[nums[i]] = &i3
	}
	return nums
}
func main() {
	b := [5]int{1, 3, 5, 8, 7}
	myTest(b, 8)
}
