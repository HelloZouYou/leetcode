// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
package main

import (
	"fmt"
	"math"
)

// 删除排序数组中的重复项
// 双指针
func removeDuplicates(nums []int) int {
	// 先处理边界
	l := len(nums)
	if l == 0 {
		return 0
	}
	left := 0

	// 双指针
	// 左指针永远指向已经确定的最后一个不重复项，并原地等待
	// 右指针不停去找下一个不重复项，找到后送到左指针面前，并让左指针右移继续等待
	for right := 1; right < l; right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

// 买卖股票的最佳时机 II
// 贪心算法
func maxProfit(prices []int) int {
	// 处理边界，不足两天的没得赚
	l := len(prices)
	if l < 2 {
		return 0
	}
	want := 0
	// 最大利润即是每个上涨的日子都完成交易，每个下跌的日子都能避开
	// 因此只要某天价格比前天大，就加到利润want中
	for i := 1; i < l; i++ {
		if prices[i] > prices[i-1] {
			want += prices[i] - prices[i-1]
		}
	}
	return want
	// 处理边界
	// l := len(prices)
	// if l < 2 {
	// 	return 0
	// }
	// 与上不同，没有直接加上每天的收益，而是用max存储阶段的最大收益，最后统一结算
	// min, max, want := prices[0], prices[0], 0
	// for i := 1; i < l; i++ {
	// 	if prices[i] < max { // 找到一个比记录max小的，则进行一次结算
	// 		want += max - min
	// 		min, max = prices[i], prices[i]
	// 	} else {
	// 		max = prices[i]
	// 	}
	// }
	// return want + max - min
}

// 旋转数组
func rotate(nums []int, k int) {
	// 处理边界
	// l := len(nums)
	// k = k % l // 移动个数超过数组长度，取余
	// if k == 0 {
	// 	return
	// }
	// 直接截取前后部分调换
	// i := l - k
	// tmp := append([]int{}, nums[0:i]...)
	// nums = append(nums[:0], nums[i:]...)
	// nums = append(nums, tmp...)

	// 处理边界，余数判断实际是否需要移动
	// l := len(nums)
	// if k%l == 0 {
	// 	return
	// }
	// 依次跳跃移动，每移动一个，用tmp记录被挤出来的值，并为其找要移动到的位置
	// 遇到循环，即已经处理过的坑位，则用visited记录，然后位置v+1找下一个循环
	// tmp, i := nums[0], 0
	// visited, v := map[int]struct{}{}, 0
	// for v < l {
	// 	i = (i + k) % l
	// 	if _, ok := visited[i]; ok {
	// 		i++
	// 		tmp = nums[i]
	// 		continue
	// 	}
	// 	nums[i], tmp = tmp, nums[i]
	// 	visited[i] = struct{}{}
	// 	v++
	// }

	// 需要先找出规律
	// 先全部翻转，再翻转前k个，再翻转后l-k个
	// l := len(nums)
	// k = k % l
	// if k == 0 {
	// 	return
	// }
	// for i, e := 0, l-1; i < e; {
	// 	nums[i], nums[e] = nums[e], nums[i]
	// 	i++
	// 	e = l - i - 1
	// }
	// for i, e := 0, k-1; i < e; {
	// 	nums[i], nums[e] = nums[e], nums[i]
	// 	i++
	// 	e = k - i - 1
	// }
	// for i, e := k, l-1; i < e; {
	// 	nums[i], nums[e] = nums[e], nums[i]
	// 	i++
	// 	e = l - i - 1 + k
	// }

	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}
	nums = append(nums[:0], append(nums[l-k:], nums[:l-k]...)...)
}

// 存在重复元素
func containsDuplicate(nums []int) bool {
	// 全部遍历，用临时map记录，发现已被记录过的即为重复元素
	v := map[int]struct{}{}
	for _, i := range nums {
		if _, ok := v[i]; ok {
			return true
		}
		v[i] = struct{}{}
	}
	return false
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	// 处理边界
	l := len(nums)
	if l == 0 {
		return 0
	}
	if l <= 2 {
		return nums[0]
	}
	// 根据异或的规律，一个数异或自己为0
	// 0异或一个数则为该数
	// 因此所有元素异或后的值，则一定为结果
	var res int
	for i := 0; i < l; i++ {
		res ^= nums[i]
	}
	return res
}

// 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	// 处理边界
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 == 0 || l2 == 0 {
		return []int{}
	}

	// 先记录a数组所有元素的出现次数，用于给b的元素一一对应
	tmp := make(map[int]int)
	for _, item := range nums1 {
		tmp[item]++
	}

	// 然后遍历b数组，找不到的直接原地删掉
	// 每找到一个存在的元素，将其对应个数减一
	for i := 0; i < l2; i++ {
		if v, ok := tmp[nums2[i]]; ok && v > 0 {
			tmp[nums2[i]]--
		} else {
			nums2 = append(nums2[:i], nums2[i+1:]...)
			l2--
			i--
		}
	}
	return nums2
}

// 有效的数独
func isValidSudoku(board [][]byte) bool {
	// 处理边界
	l := len(board)
	if l != 9 {
		return false
	}
	// 用三个map分别存储横向、竖向、3*3
	h := map[int]map[byte]struct{}{}
	s := map[int]map[byte]struct{}{}
	q := map[int]map[byte]struct{}{}
	// 全部初始化好
	for i := 0; i < l; i++ {
		h[i] = make(map[byte]struct{})
		s[i] = make(map[byte]struct{})
		q[i] = make(map[byte]struct{})
	}
	for i, g := range board { // i为行
		for k, v := range g { // k为列， v为值
			if _, ok := h[i][v]; v != '.' && ok { // 判断数字在该行是否已经出现过，出现过则无效
				return false
			}
			h[i][v] = struct{}{}
			if _, ok := s[k][v]; v != '.' && ok { // 判断数字在该列是否已经出现过，出现过则无效
				return false
			}
			s[k][v] = struct{}{}
			if _, ok := q[(i/3)*3+k/3][v]; v != '.' && ok { // 判断数字是否在当前所属的3*3出现过
				return false
			}
			// i/3*3+k/3 用于确定处于哪个3*3
			q[(i/3)*3+k/3][v] = struct{}{}
		}
	}
	return true
}

// 加一
func plusOne(digits []int) []int {
	l := len(digits)
	if l == 0 {
		return []int{}
	}

	// 从最后面的数依次判断，tmp记录是否需要进1
	tmp := 1
	for i := l - 1; tmp != 0 && i >= 0; i-- {
		digits[i] += tmp
		tmp = digits[i] / 10
		digits[i] %= 10
	}
	// 最后补充判断是否还需要进1
	if tmp != 0 {
		res := make([]int, 0, l+1)
		digits = append(append(res, tmp), digits...)
	}
	return digits
}

// 移动零
func moveZeroes(nums []int) {
	l := len(nums)
	if l == 1 {
		return
	}
	// 遍历判断，存在0则原地截取后，末位补0
	for k, i := 0, 0; i < l; i++ {
		if nums[k] == 0 {
			nums = append(nums[:k], nums[k+1:]...)
			nums = append(nums, 0)
			k-- // 截断后需要保持指针不动
		}
		k++
	}
}

// 两数之和
func twoSum(nums []int, target int) []int {
	var (
		tmp = map[int]int{}
	)
	// 遍历判断，用tmp留下自己的目标值，以及自己的位置k
	for k, v := range nums {
		if key, ok := tmp[v]; ok { // 找到目标值v，则返回目标值的位置key和自己的位置k
			return []int{key, k}
		}
		tmp[target-v] = k // 用目标数记录自己的位置k
	}
	return []int{}
}

// 旋转图像
func rotate1(matrix [][]int) {
	l := len(matrix)
	if l <= 1 {
		return
	}
	// 先上下交换
	for i := 0; i < l/2; i++ {
		matrix[i], matrix[l-i-1] = matrix[l-i-1], matrix[i]
	}
	// 再对角线交换
	for i := 1; i < l; i++ {
		for k := 0; k < i; k++ {
			matrix[i][k], matrix[k][i] = matrix[k][i], matrix[i][k]
		}
	}
	fmt.Println(matrix)
}

// 反转字符串
func reverseString(s []byte) {
	// 处理边界
	l := len(s)
	if l <= 1 {
		return
	}
	// 对半处理，头尾交换即可
	k := l / 2
	for i := 0; i < k; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

// 整数反转
func reverse(x int) int {
	// 个位数直接返回
	if x < 10 && x > -10 {
		return x
	}
	res := 0
	for x != 0 {
		s := x % 10                                   // 除10取商，即可依次拿到每个数，类似将底下的积木一个个抽到上面
		m := res*10 + s                               // 将之前低位的数乘以10拔高，再加上当前的高位数字，实现反转
		if m >= math.MaxInt32 || m <= math.MinInt32 { // 超过限制的直接返回0
			return 0
		}
		x, res = x/10, m
	}
	return res
}

// 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	// 假设只有26个字母，直接用0到25存储
	a := [26]int{}
	// 第一次遍历，记录每个字母的出现次数
	for _, v := range s {
		a[v-'a']++
	}
	// 二次遍历，找出哪个数字只出现一次
	for k, v := range s {
		if n := a[v-'a']; n == 1 {
			return k
		}
	}
	return -1
}

// 有效的字母异位词
func isAnagram(s string, t string) bool {
	// 处理边界，有效异位词起码长度要一致
	l1 := len(s)
	l2 := len(t)
	if l1 != l2 {
		return false
	}
	a := map[rune]int{}
	// 先记录第一个字符串每个字符出现的次数
	for _, v := range s {
		a[v]++
	}
	// 遍历第二个字符串，将字符一一对应减一，如果出现某个字符不够减，则说明不是有效异位词
	// 因为长度一定相等，若某个字符个数多了，则一定有某个字符个数少了
	for _, v := range t {
		if g, ok := a[v]; !ok || g == 0 {
			return false
		}
		a[v]--
	}
	return true
}

// 验证回文串
func isPalindrome(s string) bool {
	l := len(s)
	if l < 2 {
		return true
	}
	t, e := 0, l-1 // 双指针，分别指向头尾
	r := []rune(s)
	for t != e {
		// 遇到大写先全部转小写
		if r[t] >= 97 && r[t] <= 122 {
			r[t] -= 32
		}
		if r[e] >= 97 && r[e] <= 122 {
			r[e] -= 32
		}
		// 遇到非数字或字母，则直接移动指针
		if (r[t] < 48 || r[t] > 57) && (r[t] < 65 || r[t] > 90) {
			t++
			continue
		}
		if (r[e] < 48 || r[e] > 57) && (r[e] < 65 || r[e] > 90) {
			e--
			continue
		}
		// 对于数字和字母，则终于开始判断头尾是否一致，但凡遇到不一致的，直接false
		if r[t] != r[e] {
			return false
		}
		// 指针相遇时停止遍历
		if e-t == 1 {
			break
		}
		t++
		e--
	}
	return true
}

// 字符串转换整数 (atoi)
func myAtoi(s string) int {
	res := 0
	prefix := false // 用于记录是否已经找到前缀符号
	symbol := true  // 用于判断是正数还是负数
	for _, v := range s {
		g := int(v) - 48 // ascii转十进制
		if !prefix {     // 还没找到前缀符号的情况
			if v == ' ' { // 先忽略所有空格
				continue
			}
			if v == '-' { // 负号是认为找到了，同时标记为负数
				symbol = false
			} else if v == '+' { // +号也认为找到了
			} else if v >= 48 && v <= 57 { // 直接出现数字，也是认为找到了，同时还别忘了记录下该数字
				res = res*10 + g
			} else {
				return 0
			}
			prefix = true
		} else { // 符号出现后
			if v >= 48 && v <= 57 { // 将有效数字依次加上
				res = res*10 + g
			} else { // 非数字，直接over结束
				break
			}
		}
		// 每次遍历需要判断是否已经加到上限
		if res > math.MaxInt32 {
			res = math.MaxInt32 + 1 // 因为有符号数的最大值的绝对值，比最小值的绝对值小1，因此为避免如果负数，则最好+1
			break
		}
	}
	// 处理是否负数
	if !symbol {
		res = -res
	}

	// 比最小值还小则直接返回最小值
	if res < math.MinInt32 {
		res = math.MinInt32
	} else if res > math.MaxInt32 { // 而若比最大值还大，则仅返回最大值，这里化解了前面的加1
		res = math.MaxInt32
	}
	return res
}

// 实现 strStr()
func strStr(haystack string, needle string) int {
	// 空串直接返回0，符合说明
	if needle == "" {
		return 0
	}
	// 待查询的更长，肯定找不到
	l1 := len(haystack)
	l2 := len(needle)
	if l2 > l1 {
		return -1
	}
	// 长度一致则可以直接比较是否相等
	if l1 == l2 {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}
	// 双指针
	// i记录主查询的主串指针位置
	// j记录主查询的子串指针位置
	// n1记录子查询的主串指针位置
	// n2记录子查询的子串指针位置
	// f1记录主查询是否有找到第一个字符匹配
	// f2记录子查询是否有找到第一个字符匹配
	i, j, n1, n2, res, f1, f2 := 0, 0, 0, 0, 0, false, false
	for i < l1 && j < l2 {
		if haystack[i] == needle[j] { // 主查询判断字符是否一致
			if i == n1 { // 主查询和子查询的主串指针位置一致时，先不进行子查询，等下一次
				n1++
			} else {
				// 进入到这，i就会一直走在n1前面，不会再到上面的if
				if haystack[n1] == needle[n2] { // 子查询判断字符是否一致
					n1++ // 找到则子查询指针右移，一旦没有找到，n1和n2原地停止不再右移
					n2++
					if !f2 { // 一旦找到第一个匹配则标记上
						f2 = true
					}
				}
			}
			if !f1 { // 主查询标记是否找到第一个匹配
				res, f1 = i, true // res记录主查询找到第一个匹配的位置
			}
			i++
			j++
		} else { // 主查询终于匹配失败了
			if !f2 { // 如若子查询还没有开始过，则其自行计算新的i、j位置，并重置子查询的起点
				i = i - j + 1
				j, n1, n2 = 0, i, 0
			} else { // 如果子查询开始过，则从子查询的最终匹配停止的点重新开始，相当于接力
				i, j, f2 = n1, n2, false
			}
			res = i - j
		}
	}
	if !f1 || i-res < l2 {
		return -1
	}
	return res
}

func strStr1(haystack string, needle string) int {
	// 空串直接返回0，符合说明
	if needle == "" {
		return 0
	}
	// 待查询的更长，肯定找不到
	l1 := len(haystack)
	l2 := len(needle)
	if l2 > l1 {
		return -1
	}
	// 长度一致则可以直接比较是否相等
	if l1 == l2 {
		if haystack == needle {
			return 0
		} else {
			return -1
		}
	}
	for k, _ := range haystack {
		if l1-k < l2 {
			return -1
		}
		if haystack[k:l2+k] == needle {
			return k
		}
	}
	return -1
}

// 外观数列
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
}

func main() {
	fmt.Println(strStr1("aaaabb", "bb"))
}
