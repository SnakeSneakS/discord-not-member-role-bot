package core

func CountSameString(array1 []string, array2 []string) int {
	sum := 0
	for _, s1 := range array1 {
		for _, s2 := range array2 {
			if s1 == s2 {
				sum++
				break
			}
		}
	}
	return sum
}

//可変長
