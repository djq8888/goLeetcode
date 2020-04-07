/*两数之和*/
package main

func twoSum(nums []int, target int) []int {
	exist := make(map[int]int)
	for seq, n := range nums {
		exist[n] = seq
	}
	res := make([]int, 0)
	for seq, n := range nums {
		if e, ok := exist[target-n]; ok && e != seq {
			res = append(res, seq)
			res = append(res, e)
			break
		}
	}
	return res
}