// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x26epd/
package main

import (
	"fmt"
	"math/rand"
)

// 打乱数组
type Solution struct {
	ori  []int
	data []int
	l    int
}

func Constructor(nums []int) Solution {
	l := len(nums)
	res := Solution{
		ori:  nums,
		data: make([]int, l),
		l:    l,
	}
	copy(res.data, nums)
	return res
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.ori
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	for i := 0; i < this.l; i++ {
		k := rand.Intn(this.l)
		this.data[i], this.data[k] = this.data[k], this.data[i]

	}
	return this.data
}

func main() {
	data := Constructor([]int{-6, 10, 184})
	fmt.Println(data.Reset())
	fmt.Println(data.Shuffle())
	fmt.Println(data.Shuffle())
	fmt.Println(data.Shuffle())
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
