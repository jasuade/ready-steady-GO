package raindrops

import "testing"

func TestOfRaindrops(t *testing.T) {
	//Test in parallel
	//Create an structure for all test
	testCase := []struct {
		desc string
		in   int
		want string
	}{ //Fill the content of that structure with the several test
		{desc: "1. Return 'Pling' when receiving a number factor of 3", in: 9, want: "Pling"},
		{desc: "2. Return 'Plang' when receiving a number factor of 5", in: 10, want: "Plang"},
		{desc: "3. Return 'Plong' when receiving a number factor of 7", in: 28, want: "Plong"},
		{desc: "4. Return the number when receiving a number without factor", in: 34, want: "34"},
		{desc: "5. Return 'PlingPlong' when receiving a number factor of 3 and 7", in: 21, want: "PlingPlong"},
		{desc: "6. Return 'PlingPlang' when receiving a number factor of 3 and 5", in: 30, want: "PlingPlang"},
		{desc: "7. Return 'PlingPlangPlong' when receiving a number factor of 3, 5 and 7", in: 210, want: "PlingPlangPlong"},
	}
	//Create a loop to iterate among test
	for _, tcase := range testCase {
		tcase := tcase                         //Instanciate the test in a local variable to run them parrallel
		t.Run(tcase.desc, func(t *testing.T) { //Run the instanciated test
			t.Parallel()               //Run it in parallel
			got := raindrops(tcase.in) //Create result
			if got != tcase.want {     //Check result
				t.Errorf("Not correct result received, but %s instead.", got) //Throw an error if the result is not the expected
			}
		})
	}

	// t.Run("Return 'Pling' when receiving a number factor of 3", func(t *testing.T) {
	// 	got := raindrops(9)
	// 	want := "Pling"
	// 	if got != want {
	// 		t.Errorf("Not 'Pling' received, but %s instead.", got)
	// 	}

	// })
	// t.Run("Return 'Plang' when receiving a number factor of 5", func(t *testing.T) {
	// 	got := raindrops(10)
	// 	want := "Plang"
	// 	if got != want {
	// 		t.Errorf("Not 'Plang' received, but %s instead.", got)
	// 	}

	// })
	// t.Run("Return 'Plong' when receiving a number factor of 7", func(t *testing.T) {
	// 	got := raindrops(28)
	// 	want := "Plong"
	// 	if got != want {
	// 		t.Errorf("Not 'Plong' received, but %s instead.", got)
	// 	}

	// })
	// t.Run("Return the number when receiving a number without factor", func(t *testing.T) {
	// 	got := raindrops(34)
	// 	want := "34"
	// 	if got != want {
	// 		t.Errorf("Not digit number received, but %s instead.", got)
	// 	}

	// })
	// t.Run("Return 'PlingPlong' when receiving a number factor of 3 and 7", func(t *testing.T) {
	// 	got := raindrops(21)
	// 	want := "PlingPlong"
	// 	if got != want {
	// 		t.Errorf("Not 'PlingPlong' received, but %s instead.", got)
	// 	}

	// })
	// t.Run("Return 'PlingPlang' when receiving a number factor of 3 and 5", func(t *testing.T) {
	// 	got := raindrops(30)
	// 	want := "PlingPlang"
	// 	if got != want {
	// 		t.Errorf("Not 'PlingPlang' received, but %s instead.", got)
	// 	}

	// })
	// t.Run("Return 'PlingPlangPlong' when receiving a number factor of 3, 5 and 7", func(t *testing.T) {
	// 	got := raindrops(210)
	// 	want := "PlingPlangPlong"
	// 	if got != want {
	// 		t.Errorf("Not 'PlingPlangPlong' received, but %s instead.", got)
	// 	}

	// })

}
