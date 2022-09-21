package calPoints

import (
	"fmt"
	"strconv"
)

func main() {
	s := []string{"5", "2", "C", "D", "+"}
	fmt.Println(calPoints(s))
}
func calPoints(operations []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(operations); i++ {
		if operations[i] == "C" {
			stack = stack[:len(stack)-1]
		} else if operations[i] == "D" {
			peekNum := stack[len(stack)-1]
			peekNum = peekNum * 2
			stack = append(stack, peekNum)
		} else if operations[i] == "+" {
			firstNum := stack[len(stack)-1]
			secondNum := stack[len(stack)-2]
			tempNum := firstNum + secondNum
			stack = append(stack, tempNum)
		} else {
			peekNum, _ := strconv.Atoi(operations[i])
			stack = append(stack, peekNum)
		}
	}
	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}
	return sum
}
