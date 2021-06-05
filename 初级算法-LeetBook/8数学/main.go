// https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2kxrh/
package main

import (
	"fmt"
	"strconv"
)

// Fizz Buzz
func fizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

// 计数质数
func countPrimes(n int) int {
	tmp, nums := make([]bool, n), 0
	for i := 2; i < n; i++ {
		if !tmp[i] {
			nums++
			for j := i + i; j < n; j += i {
				tmp[j] = true
			}
		}
	}
	return nums
}

// 3的幂
func isPowerOfThree(n int) bool {
	// return n > 0 && (n == 1 || n%3 == 0 && isPowerOfThree(n/3))
	for n > 1 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// 罗马数字转整数
func romanToInt(s string) int {
	res, tmp := 0, 0
	dic := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	for _, v := range s {
		res += dic[v]
		if dic[v] > tmp {
			res -= tmp + tmp
		}
		tmp = dic[v]
	}
	return res
}

func main() {
	fmt.Println(romanToInt("IM"))
}
