package GenerateDocument

/*
遍历map
for k, v := range docMap {

	}
*/

func GenerateDocument(characters string, document string) bool {
	// Write your code here.
	dictMap := make(map[rune]int, 0)
	runeArr := []rune(characters)
	for i := 0; i < len(characters); i++ {
		if _, ok := dictMap[runeArr[i]]; ok {
			dictMap[runeArr[i]]++
		} else {
			dictMap[runeArr[i]] = 1
		}
	}

	docMap := make(map[rune]int, 0)
	runeArr2 := []rune(document)
	for i := 0; i < len(runeArr2); i++ {
		if _, ok := docMap[runeArr2[i]]; ok {
			docMap[runeArr2[i]]++
		} else {
			docMap[runeArr2[i]] = 1
		}
	}

	for k, v := range docMap {
		if value, ok := dictMap[k]; ok && v <= value {
			continue
		} else {
			return false
		}
	}

	return true
}
