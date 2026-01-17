//  
/*
question:
1. variable kii?
ans :variable holo akta container jekhane data rakha hoi.
2. akta computer er ki ki thake ?
ans: cpu,ram,hardisk
*/
/**/
package main
import "fmt"

func main(){
	// var a = 10
	// var a int = 10
	a := 10

// a := 20 aii vabe likha jabe nah jokhon kono variable er value re assign kori tokhon : use kora jabe nah :use hobe first time use er somoi 
a = 20 
// a= "dipto" ai vabe kora jabe nah karon first bar jokhon value assign korsi tokhon value inteser type chilo so go auto type buje felse tai type mismatch kora jabe nah.

	fmt.Println(a)

	const p = 100

	// p = 5 aita kora jabe nah const variable er man change kora jai nah
	fmt.Println(p)
}