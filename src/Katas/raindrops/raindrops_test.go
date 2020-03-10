package raindrops

import "testing"

func TestOfRaindrops(t *testing.T) {
	t.Run("Return 'Pling' when receiving a number factor of 3", func(t *testing.T) {
		got := raindrops(9)
		want := "Pling"
		if got != want {
			t.Errorf("Not 'Pling' received, but %s instead.", got)
		}

	})
	t.Run("Return 'Plang' when receiving a number factor of 5", func(t *testing.T) {
		got := raindrops(10)
		want := "Plang"
		if got != want {
			t.Errorf("Not 'Plang' received, but %s instead.", got)
		}

	})
	t.Run("Return 'Plong' when receiving a number factor of 7", func(t *testing.T) {
		got := raindrops(28)
		want := "Plong"
		if got != want {
			t.Errorf("Not 'Plong' received, but %s instead.", got)
		}

	})
	t.Run("Return the number when receiving a number without factor", func(t *testing.T) {
		got := raindrops(34)
		want := "34"
		if got != want {
			t.Errorf("Not digit number received, but %s instead.", got)
		}

	})
	t.Run("Return 'PlingPlong' when receiving a number factor of 3 and 7", func(t *testing.T) {
		got := raindrops(21)
		want := "PlingPlong"
		if got != want {
			t.Errorf("Not 'PlingPlong' received, but %s instead.", got)
		}

	})
	t.Run("Return 'PlingPlang' when receiving a number factor of 3 and 5", func(t *testing.T) {
		got := raindrops(30)
		want := "PlingPlang"
		if got != want {
			t.Errorf("Not 'PlingPlang' received, but %s instead.", got)
		}

	})
	t.Run("Return 'PlingPlangPlong' when receiving a number factor of 3, 5 and 7", func(t *testing.T) {
		got := raindrops(210)
		want := "PlingPlangPlong"
		if got != want {
			t.Errorf("Not 'PlingPlangPlong' received, but %s instead.", got)
		}

	})

}
