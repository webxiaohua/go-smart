/*
125.验证回文串
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个回文串。
字母和数字都属于字母数字字符。
给你一个字符串 s，如果它是回文串，返回 true ；否则，返回 false 。

示例 1：
输入: "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。

示例 2：
输入："race a car"
输出：false
解释："raceacar" 不是回文串。

提示：
1 <= s.length <= 2 * 105
s 仅由可打印的 ASCII 字符组成
*/
package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func main(){
	s := "race a car"
	res := isPalindrome2(s)
	fmt.Println(res)
}

// 循环法
func isPalindrome(s string) bool {
	// 移除非数字&字母后的字符串
	var buffer bytes.Buffer
	lower := strings.ToLower(s)
	for _,c := range lower {
		if (c >= '0' && c <= '9') || (c >= 'a' && c<='z') || (c >= 'A' && c<='Z') {
			buffer.WriteString(string(c))
		}
	}
	newS := buffer.String()
	length := len(newS)
	if length == 1 {
		return true
	}
	middle := length / 2
	for i:=0;i<middle;i++ {
		if newS[i] != newS[length-i-1] {
			return false
		}
	}
	return true
}

// 双指针
func isPalindrome2(s string) bool{
	reg,_ :=regexp.Compile("[^a-zA-Z0-9]")
	s = reg.ReplaceAllString(strings.ToLower(s),"")
	left:=0
	right := len(s) - 1
	for {
		if left >= right {
			return true
		}
		if s[left] == s[right]{
			left++
			right--
		}else{
			return false
		}
	}
}