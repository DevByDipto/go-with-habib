/*
1. go runtime kii ?
2. os scheduler kii ?
3. go runtime scheduler vs go routine scheduler kii ?
4. thread er stack koto go routine er stack koto ?
ans: thread er stack 8 mb and go routine er stack 2 kb
5. ->go routine/mini thread/vartual thread/ logical thread 
6. ->program terminate mane program close hoi jawa


*/
// package main

// import (
// 	"fmt"
// 	"time"
// )

// const p = 11

// func add(a, b int) {
// 	fmt.Println(a + b)
// }

// func printHello(num int) {
// 	fmt.Println("Hello Habib", num)
// 	add(2, 4)
// }

// func main() {
// 	// Launching goroutines
// 	go printHello(1)
// 	go printHello(2)
// 	go printHello(3)
// 	go printHello(4)
// 	go printHello(5)


// 	a := 10 
// 	fmt.Println(a, " ", p)

	
// 	time.Sleep(5 * time.Second)
// }
// run the code sea the answer and simulate it.
// ------------------------------------------------------------------
package main

import (
	"fmt"
	"time"
)

const p = 11



func printHello(num int) {
	// time.Sleep(4 * time.Second) 
	// or 
	time.Sleep(5 * time.Second)
	fmt.Println("Hello Habib", num)

}

func main() {
	// Launching goroutines
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)

fmt.Println("hello")
	
}
// run the code sea the answer and simulate it.