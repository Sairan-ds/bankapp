package deposit

import "fmt"

func ExampleCalculate() {
	fmt.Println(Calculate(0, "TJS"))
	fmt.Println(Calculate(0, "USD"))
	fmt.Println(Calculate(50000000, "TJS"))
	fmt.Println(Calculate(50000000, "USD"))
	fmt.Println(Calculate(100000000, "TJS"))
	fmt.Println(Calculate(100000000, "USD"))

	// Output:
	// 0 0
	// 0 0
	// 166666 250000
	// 41666 83333
	// 333333 500000
	// 83333 166666

}
