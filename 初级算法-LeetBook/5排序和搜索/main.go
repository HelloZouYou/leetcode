// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2322r/
package main

import "fmt"

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	t1, t2 := 0, 0
	nums1 = nums1[:m]
	// 双指针移动
	for m != 0 && t1 < len(nums1) && t1 < m+n && t2 < n {
		if nums1[t1] >= nums2[t2] {
			// nums1 = append(nums1[:t1], append([]int{nums2[t2]}, nums1[t1:]...)...)
			nums1 = append(nums1, nums2[t2])
			l := len(nums1)
			nums1 = append(nums1[:t1], append(nums1[l-1:], nums1[t1:l-1]...)...)
			t2++
		}
		t1++
	}
	nums1 = append(nums1, nums2[t2:]...)
	fmt.Println(nums1)
}

// 第一个错误的版本
func firstBadVersion(n int) int {
	min, max := 1, n
	// 二分法
	for max-min > 1 {
		mid := (min + max) / 2
		if isBadVersion(mid) {
			max = mid
		} else {
			min = mid + 1
		}
	}
	// 最后不确定的只有min，max一定是错的
	if isBadVersion(min) {
		return min
	}
	return max
}

func isBadVersion(n int) bool {
	if n >= 5 {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstBadVersion(10))
}
