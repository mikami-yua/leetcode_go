package permuteUnique

//leetcode 47 全排列II
func permuteUnique(nums []int) [][]int {
	permutions := make([][]int, 0)
	helpGet(nums, make([]int, 0), &permutions)

	/*
		没有去重
		1.先再algoexpert中跑一边
		2.加入去重的功能再leetcode中run
	*/
	return nil
}

/*
当需要修改总返回值时，类似二重的list，使用*的方式定义函数

使用copy函数不影响内存地址

slice删除指定下标索引的文件
//删除元素 3 索引(下标) 3
index := 3
// 这里通过 append 方法 分成两个然后合并
// append(切片名,追加的元素) 切片名这里我们进行切割一个新的切片DelIndex[:index] 追加的元素将索引后面的元素追加
// DelIndex[index+1:]...) 为什么追加会有...三个点, 因为是一个切片 所以需要展开
DelIndex = append(DelIndex[:index], DelIndex[index+1:]...)

当appen的slice是同类型时，才需要...，次一级的append直接向后加即可
*/

func help(array []int, curPermution []int, permutions *[][]int) {
	if len(array) == 0 && len(curPermution) > 0 {
		*permutions = append(*permutions, curPermution)
	} else {
		for i := 0; i < len(array); i++ {
			newArray := make([]int, len(array))
			copy(newArray, array)
			newArray = newArray[:len(newArray)-1]
			newPermutation := make([]int, 0)
			newPermutation = append(newPermutation, array[i])
			help(newArray, newPermutation, permutions)
		}
	}
}

func helpGet(restArray []int, curPermution []int, permutions *[][]int) {
	if len(restArray) == 0 {
		*permutions = append(*permutions, curPermution)
	} else {
		for i := 0; i < len(restArray); i++ {
			newArray := make([]int, len(restArray))
			copy(newArray, restArray)
			newArray = append(newArray[:i], newArray[i+1:]...)
			newPermutation := make([]int, 0)
			newPermutation = append(newPermutation, newArray[i])
			help(newArray, newPermutation, permutions)
		}
	}
}
