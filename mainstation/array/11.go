package array

import (
	"math"
)

func maxArea(height []int) int {
	// 头尾双指针

	/*
		先将头尾指针作为两块板分别放在容器的最左边和最后边，然后逐渐向内移动这两条板来试图获取最大的容量
		该题参考木桶原理，木桶盛水的多少只与最短的板有关
		在头尾指针向内移动的同时，容器的底长在减少，因此分以下几种情况：
		1.向内移动长板，短板不动：
			1.1 如果长板变长，容器的高只与短板有关，则移动后容器的高还是之前的短板，而容器的底随着长板的向内移动减少了，所以这次移动之后必定比移动之前容器的容量小
			1.2 如果长板变短，容器的高只与短板有关，则移动后容器的高可能是之前的短板，甚至可能是长板移动后变成的比短板更短的板，而容器的底随着长板的向内移动减少了，所以这次移动之后必定比移动之前容器的容量小
			1.3 总结：长板向内移动，不管是变长还是变短，移动之后的容器必定比移动之前容量小
		2.向内移动短板，长板不动：
			2.1 如果短板变长，容器的高只与短板有关，则移动后容器的高可能是短板移动后变成的更长的板，甚至可能是之前的长板，而容器的底随着短板的向内移动减少了，但容器的高变高了，因此这次移动可能会导致容量增加、不变或者减少
			2.2 如果短板变短，容器的高只与短板有关，则移动后容器的高是短板移动后更短的板，而容器的底随着短板的向内移动减少了，所以这次移动之后必定比移动之前容器的容量小
			2.3 总结：短板向内移动，则容器的容量可能变大、不变或变小
		因此应该向内移动短板，来尝试获取比移动前更大的容器容量
	*/

	// 1.先将头尾指针作为两块板分别放在容器的最左边和最后边
	left := 0
	right := len(height) - 1

	// 2.初始化容量为最小值
	res := math.MinInt

	// 3.逐渐向内移动这两条板来试图获取最大的容量
	for left < right {
		// 3.1 获取当前的容量
		temp := (right - left) * minInt(height[left], height[right])
		// 3.2 比较移动后的容量和移动前的容量哪个大，选择较大的赋值给结果
		res = maxInt(res, temp)
		// 3.3 比较左右板哪个更短，将更短的向内移动，来试图寻找更大的容量
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	// 4.将结果返回
	return res
}
