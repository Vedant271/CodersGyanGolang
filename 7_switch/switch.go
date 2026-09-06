package main

import (
	"fmt"
)

func main() {
	// simple switch
	// i := 4
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }

	// day := time.Now().Weekday()
	// switch day {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Working day")
	// }

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case string:
			fmt.Println("Input is String")
		case int:
			fmt.Println("Input is Integer")
		case bool:
			fmt.Println("Input is Boolean")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI(50.0)
}
