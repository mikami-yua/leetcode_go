package compressString

import "strconv"

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}
	runeArray := []rune(S)
	ans := ""
	curRune := runeArray[0]
	tempCount := 1
	for i := 1; i < len(runeArray); i++ {
		if runeArray[i] == curRune {
			tempCount++
		} else {
			ans = ans + string(curRune)
			ans = ans + strconv.Itoa(tempCount)
			curRune = runeArray[i]
			tempCount = 1
		}
	}
	ans = ans + string(curRune)
	ans = ans + strconv.Itoa(tempCount)
	return ans
}
