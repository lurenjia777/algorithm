package main

// 快速排序
func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := nums[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for _, num := range nums[1:] {
		if num <= mid {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	ret := QuickSort(left)
	ret = append(ret, mid)
	ret = append(ret, QuickSort(right)...)

	return ret
}
