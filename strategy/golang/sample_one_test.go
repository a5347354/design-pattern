package golang

func ExampleTrain() {
	user := NewUser("John", Train{})
	user.travel()
	// Output: John搭火車
}

func ExampleAirplane() {
	user := NewUser("Ada", Airplane{})
	user.travel()
	// Output:
	// Ada搭飛機
}
