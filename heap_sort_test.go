package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeapSort(t *testing.T) {

	rand.Seed(time.Now().Unix())
	nums := make([]int, 0)
	for i := 0; i < 50; i++ {
		nums = append(nums, int(rand.Int31n(10000)))
	}

	HeapSort(nums)
	fmt.Println(nums)
}
