/*
1. stack frame er synoname function fram
2. akta ram er bivinno cell e j local variable o argument value gulo rakha hoi se gulo dorkar porle kivabe khuje ber kora hoi ?
ans: register set er base pointer er adress value  sate nirdishoto value jog biyog kore ber kora hoi.

*/

package main

import "fmt"

func add(x int, y int) int {
	var res int
	res = x + y
	return res
}

func main() {
	a := 10
	sum := add(a, 4)
	fmt.Println(sum)
}
// task - simulate how it work using sp bp