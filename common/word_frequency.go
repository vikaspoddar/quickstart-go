package common

func WordFrequency(word string) map[string]int {
	var frequencyTable = make(map[string]int)
	for _, char := range word {
		if char == ' ' {
			continue
		}
		frequencyTable[string(char)]++
	}
	return frequencyTable
}
