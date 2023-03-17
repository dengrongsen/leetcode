/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */
package leetcode

import (
	"math"
	"strconv"


)

// @lc code=start
func reverse(x int) int {
	var flag bool
	if x < 0 {
		flag = true
		x = ^x + 1
	}

	xs := strconv.FormatInt(int64(x),10)
	
	xsr := []rune(xs)
	i,j := 0,len(xsr)-1
	for ; i<j; {
		xsr[i],xsr[j] = xsr[j],xsr[i]
		i++;j--
	}

	res, err := strconv.Atoi(string(xsr))
	if err != nil {
		return 0
	}

	if flag {
		res = -res
	}

	// 校验
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}
// @lc code=end

// @lc code=start
func reverse2(x int) int {
	res := 0
	for ; x != 0 ; {
		res = res * 10 + x % 10
		x /= 10
	}

	// 校验
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}

	return res
}
// @lc code=end

