package isPalindrome

/*
字符串专为rune数组的方式
[]rune(s)
ASCII码大小写转换32,小写字母-32

*/

func isPalindrome(s string) bool {
	runeArray := []rune(s)
	charArray := make([]rune, 0)

	for i := 0; i < len(runeArray); i++ {
		if runeArray[i] <= 'Z' && runeArray[i] >= 'A' {
			charArray = append(charArray, runeArray[i]+32)
		} else if runeArray[i] <= 'z' && runeArray[i] >= 'a' {
			charArray = append(charArray, runeArray[i])
		} else if runeArray[i] <= '9' && runeArray[i] >= '0' {
			charArray = append(charArray, runeArray[i])
		} else {
			continue
		}
	}

	return isPalind(charArray)
}

// 判断回文
func isPalind(arr []rune) bool {
	leftIndex, rightIndex := 0, len(arr)-1
	for leftIndex < rightIndex {
		if arr[leftIndex] == arr[rightIndex] {
			leftIndex++
			rightIndex--
		} else {
			return false
		}
	}
	return true
}
