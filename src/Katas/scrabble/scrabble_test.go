package scrabble

import "testing"

func TestMapper(t *testing.T) {
	t.Run("GIVEN A I should get 1", func(t *testing.T) {
		// Arrange
		letter := "A"
		const score = 1
		// Act
		got := Mapper(letter)
		// Assert
		if got != score {
			t.Errorf("expected %d but received %d", score, got)
		}
	})

	t.Run("GIVEN A, E, I, O, U, L, N, R, S, or T should get 1", func(t *testing.T) {
		// Arrange
		letters := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
		score := 1
		// Act
		for _, letter := range letters {
			got := Mapper(letter)
			// Assert
			if got != score {
				t.Errorf("expected %d but received %d", score, got)
			}
		}
	})

	t.Run("GIVEN D, G should return 2", func(t *testing.T) {
		// Arrange
		letters := []string{"D", "G"}
		score := 2

		// Act
		for _, letter := range letters {
			got := Mapper(letter)
			if got != score {
				t.Errorf("expected %d but received %d", score, got)
			}
		}
	})

	t.Run("GIVEN the word 'GO' the function should return 3", func(t *testing.T) {
		got := ("GO")
		score := 3
		if got != score {
			t.Errorf("Expected result %d but received %d", score, got)
		}
	})

}
