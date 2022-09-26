package array

func minArray1(numbers []int) int {
	// 暴力找降序位置

	// 特殊情况
	var lengthOfNumbers int = len(numbers)
	if lengthOfNumbers == 1 {
		return numbers[0]
	}

	// 从头开始遍历，找到降序的地方
	for i := 0; i < lengthOfNumbers-1; i++ {
		if numbers[i+1] < numbers[i] {
			return numbers[i+1]
		}
	}

	// 如果没有降序的地方，说明没有翻转，此时返回第一个数
	return numbers[0]
}

func minArray2(numbers []int) int {
	// 二分查找

	var leftOfNumbers int = 0
	var rightOfNumbers int = len(numbers) - 1
	var mid int

	// while循环
	for leftOfNumbers <= rightOfNumbers {

		mid = (leftOfNumbers + rightOfNumbers) / 2
		if numbers[mid] > numbers[rightOfNumbers] {
			leftOfNumbers = mid + 1
		} else if numbers[mid] < numbers[rightOfNumbers] {
			rightOfNumbers = mid
		} else {
			rightOfNumbers--
		}
	}

	return numbers[leftOfNumbers]
}
