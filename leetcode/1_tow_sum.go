/*
1.两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值 target  的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
样例
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
*/
package main

import "fmt"

func main (){
	nums := []int{2,7,11,15}
	target := 9
	res := twoSum3(nums,target)
	fmt.Println(res)
}

// 暴力破解法，双重循环，时间复杂度 O(N^2)
func twoSum1(nums []int,target int) (res []int) {
	if len(nums) < 2 {
		return
	}
	res = make([]int,0)
	for i:=0;i<len(nums);i++ {
		for j:=1;j<len(nums);j++ {
			if nums[i] + nums[j] == target {
				res = append(res,i,j)
				return
			}
		}
	}
	return
}

// 哈希法，遍历一次，时间复杂度 O(N)，需要额外的 O(N) 空间
func twoSum2(nums []int,target int)(res []int){
	res = make([]int,0)
	tmpHash := make(map[int]int)
	for i,v:=range nums {
		tmpHash[v] = i
	}
	for i:=0;i<len(nums);i++{
		if j,ok := tmpHash[target - nums[i]];ok {
			res = append(res,i,j)
			return
		}
	}
	return
}

// 哈希法精炼写法，时间复杂度 O(N)，需要额外 O(N) 的空间
func twoSum3(nums []int,target int)(res []int){
	res = make([]int,0)
	tmpHash := make(map[int]int)
	for i,v := range nums {
		sub := target - v
		if j,ok := tmpHash[sub];ok {
			res = append(res,i,j)
			return
		}
		tmpHash[v] = i
	}
	return
}