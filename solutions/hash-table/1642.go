package hash_table

func MaxLengthBetweenEqualCharacters(s string) int {
	mapLastSeenChar := make(map[rune]int, 0)
	mapLengthEachChar := make(map[rune]int, 0)
	maxLength := 0

	for i, char := range s {
		if index, exist := mapLastSeenChar[char]; exist {
			mapLengthEachChar[char] = i - (index + 1)

			if mapLengthEachChar[char] > maxLength {
				maxLength = mapLengthEachChar[char]
			}

			continue
		}

		mapLastSeenChar[char] = i
	}

	if len(mapLengthEachChar) != 0 {
		return maxLength
	}

	return -1
}

/**
Example Case 2
Given s = "cabbac"
Output = 4

Loop
0, c
1, a
2, b
3, b
4, a
5, c

Found a Length = 2
keep max length = 2

Found b Length = 0
0 > 2 skip

Found c Length = 4
4 > 2 keep max length

return max length

**/
