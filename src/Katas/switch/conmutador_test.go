package conmutador

import (
	"testing"
)

//Depict with SO Go is running over
func TestSwitch(t *testing.T) {
	t.Run("It should return 'OS x' when calling the function", func(t *testing.T) {
		got := conmutador()
		want := "OS x"
		if got != want {
			t.Errorf("The result %s was not the expected", got)
		}
	})

	// t.Run("If should return 'Windows' when callling the function in a Win system", func(t *testing.T) {
	// 	got := conmutador()
	// 	want := "Windows"
	// 	if got != want {
	// 		t.Errorf("The result %s was not the expected", got)
	// 	}
	// })

}
