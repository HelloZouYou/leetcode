// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2uudv/
package main

import (
	"fmt"
	"math"
)

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
	// 直接逐位比较
	for k := range haystack {
		if l1-k < l2 { // 最后长度不够了不用再比了
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
	str := countAndSay(n - 1)
	l := len(str)
	res, s, e := "", 0, 0

	for e < l {
		for e < l && str[s] == str[e] {
			e++
		}
		res += string([]rune{rune(e - s + 48)}) + str[s:s+1]
		s = e
	}
	return res
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}
	res := strs[0]
	for k, v := range strs {
		// 忽略第一个
		if k == 0 {
			continue
		}
		// 取最小长度
		l1, l2 := len(res), len(v)
		if l2 < l1 {
			l1 = l2
		}
		// 逐渐截取比较
		for g := l1; g >= 0; g-- {
			res = res[:g]
			if res == v[:g] {
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abc", "cbd"}))
}
