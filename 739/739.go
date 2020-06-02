package main

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T), len(T))
	stack := make([]int, 0, len(T))
	for i, v := range T {
		for len(stack) != 0 && T[stack[len(stack) - 1]] < v {
			res[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for len(stack) < len(T) {
		stack = append(stack, 0)
	}
	return res
}
