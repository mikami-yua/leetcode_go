package productsum

import "fmt"

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	// Write your code here.

	return help(array, 1, 0)
}

/*
如果需要两次类型断言，可以使用if elseif else结构
*/

func help(array SpecialArray, deep int, index int) int {
	sum := 0
	for index < len(array) {
		temp := array[index]

		if value, ok := temp.(SpecialArray); ok {
			param := value
			sum += help(param, deep+1, 0)
		} else if value, ok := temp.(int); ok {
			sum += value
		} else {
			fmt.Println("error")
		}
		index++
	}
	return sum * deep
}
