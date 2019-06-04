package main

func HeapSort(nums []int) []int {

	lenth := len(nums)

	if lenth <= 1 {
		return nums
	}
	// 构建堆结构
	for i := lenth/2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, lenth)
	}

	for j := lenth - 1; j >= 0; j-- {
		swap(nums, 0, j, j+1)
		adjustHeap(nums, 0, j)
	}
	return nums
}

func swap(nums []int, i, j, lenth int) {
	if lenth <= len(nums) && i < lenth && j < lenth {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func adjustHeap(nums []int, i, lenth int) {
	temp := -1
	left := 2*i + 1
	right := 2*i + 2
	if left < lenth && right < lenth {
		if nums[left] > nums[right] {
			temp = left
		} else {
			temp = right
		}
	} else {
		temp = left
	}
	if temp == -1 || temp >= lenth {
		return
	}
	if nums[i] < nums[temp] {
		swap(nums, i, temp, lenth)
		adjustHeap(nums, temp, lenth)
	}
}
