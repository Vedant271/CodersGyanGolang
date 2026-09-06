package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// infinite loop
	// for {
	// 	fmt.Println("1")
	// }

	// classic for loop
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// break, continue
	// for i := 0; i < 5; i++ {
	// 	// if i == 2 {
	// 	// 	break
	// 	// } // 0 1
	// 	if i == 2 {
	// 		continue
	// 	} // 0 1 3 4
	// 	fmt.Println(i)
	// }

	// range
	for i := range 3 {
		fmt.Println(i)
	}
}
