// Control Flow: Loops

// For Loops
// Nested Loop Control
// Loop Control Statement Like Break, Continue and GoTo




package main

import "fmt"

func main() {	

	// 1. For Loop
//  for i := 0; i < 10; i++ {
// 	fmt.Println(i)
//  }

// 2. Nested Loop
// for i:=1; i <= 10; i++ {
// 	for j:=1; j <= 10; j++ {
// 		fmt.Printf("%d x %d = %d\t", i, j, i*j)
// 	}
// 	fmt.Println()
// }

// 3. Loop Control Statement
// 3A. Break
// for i :=0; i < 10; i++ {
// 	fmt.Println("hi there, Mr.", i)
// 	if i == 5 {
// 		break
// 	}
// }

// 3B. Continue
// for i :=0; i < 10; i++ {
// 	if i % 2 == 0 {
// 		continue
// 	}
// 	fmt.Println("hi there, Mr.", i)
// }

// 3C. GoTo
// for i :=0; i < 10; i++ {
// 	fmt.Println("hi there, Mr.", i)
// 	 if i == 5 {
// 		goto end
// 	 }
// }
// end: {
// 	fmt.Println("hi there, Mr. Endrewvaski!!")
// }

// i := 0;

// start:
// if i < 10 {
// 	fmt.Println("Hello Mr. ", i)
// 	i++
// 	goto start
// }

// 4. Infinite Loop
// for {
// 	fmt.Println("Hello World")
// }
start:
fmt.Println("This is an infinite loop!")
 goto start
}