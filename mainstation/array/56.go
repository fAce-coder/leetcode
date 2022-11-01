package array

import (
	"sort"
)

func merge55(intervals [][]int) [][]int {
	/*
		首先按照每组数的第一个数，从小到大进行排序
		如果后面一组数的第一个数大于前面一组数的第二个数：说明这两个区间不重合，直接加入结果集不进行合并
		如果后面一组数的第一个数小于等于前面一组数的第二个数：说明这两个区间重合，将两个区间合并
	*/

	// 1.初始化结果集
	res := make([][]int, 0)

	// 2.按照每组数的第一个数，从小到大进行排序
	sort.Sort(sortable2DNums(intervals))

	// 3.遍历原数组，将每组数加入或合并到结果集中
	for _, interval := range intervals {
		if len(res) == 0 || res[len(res)-1][1] < interval[0] {
			// 如果结果集为空，则直接加入结果集
			// 或者后面一组数的第一个数大于前面一组数的第二个数：说明这两个区间不重合，直接加入结果集不进行合并
			res = append(res, interval)
		} else {
			// 如果后面一组数的第一个数小于等于前面一组数的第二个数：说明这两个区间重合，将两个区间合并
			res[len(res)-1][1] = maxInt(res[len(res)-1][1], interval[1])
		}
	}

	// 4.返回结果集
	return res
}
