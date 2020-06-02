package main

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	//遍历数组，并使用单调栈存储高度
	stack := make([]int, 0, len(heights))
	max := 0
	for i, v := range heights {
		for len(stack) > 0 && v < heights[stack[len(stack) - 1]] {
			h := heights[stack[len(stack)-1]]   //栈顶元素的高度
			stack = stack[:len(stack)-1]
			w := i                              //栈顶元素可以形成矩形的最大宽度
			if len(stack) != 0 {
				w = i - stack[len(stack)-1] - 1
			}
			if max < h * w {
				max = h * w
			}
		}
		stack = append(stack, i)
	}

	//如果栈不为空，则说明栈中所有元素都可以都可以拓展到最后
	for len(stack) > 0 {
		h := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		w := len(heights)
		if len(stack) != 0 {
			w = len(heights) - stack[len(stack)-1] - 1
		}
		if max < h * w {
			max = h * w
		}
	}

	return max
}