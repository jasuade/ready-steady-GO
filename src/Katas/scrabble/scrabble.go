package scrabble

// This maps a letter to a score
func Mapper(letter string) int {
	lettersWithScore1 := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	lettersWithScore2 := []string{"D", "G"}

	if itemExists(letter, lettersWithScore1) {
		return 1
	}
	if itemExists(letter, lettersWithScore2) {
		return 2
	}
	return 0
}

func itemExists(item string, itemList []string) bool {
	for i := 0; i < len(itemList); i++ {
		if itemList[i] == item {
			return true
		}
	}
	return false
}
