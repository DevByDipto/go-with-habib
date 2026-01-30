//  package main

// import "fmt"

// func main(){
// 	arr := [6]string {"this","is","a","go","interview","question"}
// fmt.print(arr)

// s := arr[1:4] // slice nijer moddhe 3 ta value rakhe 
// // 1. pointer-> slice kora array er value gulor first value baa element er adress
// // 2. length-> slice array er moddhe koita value baa element rakhlo jmn aikhane len=3
// // 3. capacity-> jekhan theke clice kora shuru korse se value ta hoso oi original array te ar koita value baa element ase setar totall number jmn amader aikhetre cap=5

// fmt.print(s)

// }

// task simulate : fmt.print(arr) ai porjonto to parbo then main stack fram e s name akta sell nibe tate akta info rakhbe ptr,len,cap then jokhon s print hobe tokhon dekbe main er moddhe s ase kina dekbe ase then s er value dekbe ptr somthing ar len-3 so ptr holo arr er slice start value oi jaiga theke 3 ta value tule ane print korbe. 
// 

//------------------------------------------------------------------------------------

// package main

// import "fmt"

// func main(){    //    18     19   20  21       22         23
// 	arr := [6]string {"this","is","a","go","interview","question"}
// fmt.print(arr)

// s := arr[1:4] // ["is","a","go"]
// fmt.print(s)

// s1 := s[1:2] // ["a"]
// fmt.print(s1)
// fmt.print(lan(s1))
// fmt.print(cap(s1))
// }

// code simulation : fmt.print(s) ai porjonto to parbo. ar por jokhon dekha jabe j s1 := s[1,2] ache tokhon main er moddhe akta sell er nam hobe s1  then s er kase jabe ja akti slice header ar tate ache ptr-19,cap-5,len-3 so se s er ptr dhore arr er kache jabe sekhan theke 1:2 unojai data niye asbe ar tate thakbe ptr-20,len-1,cap-4 name akta slice header data thakbe.
// -----------------------------------------------------------------------------

// package main

// import "fmt"

// func main() {

//     s1 := []int{1, 2, 3}  // its call sclice litaral

//     fmt.Println(s1)
//     fmt.Println(len(s1)) // result 3
//     fmt.Println(cap(s1)) // result 3
// }

// task simulation: main porjonto parbo then main call hobe stack fram toiri hobe tate 3 ta sell nibe se sell e 123 value bosbe then s1 name akta cell alocate hobe tate ptr,cap,len slice header rakhbe ar porer ta parbo.

// -----------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {

//     s := make([]int,3)  // [0,0,0] len=3 cap=3
//     s[0]=5 // [5,0,0] len=3 cap=3

//     fmt.Println(s)
//     fmt.Println(len(s)) // result 3
//     fmt.Println(cap(s)) // result 3

// }

// -----------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {

//     s := make([]int,3,5)  // [0,0,0] len=3 cap=5
//     s[0]=5 // [5,0,0] len=3 cap=5
// s[2]=5
// s[4]=5
//     fmt.Println(s)
//     fmt.Println(len(s)) // result 3
//     fmt.Println(cap(s)) // result 5
// }

// task simulation :  s := make([]int,3,5) ai porjonto parbo then s[0]=5 aii line e jokhon asbe tohon dekbe main er vitor s ase kina then s pabe then s er slice header theke pointer er maddhome array er adress e pochabe ja adress er first element er adress then adress er shate 0 jog korbe then j adress pabe tar value 5 kore dibe then same vabe  dekbe main er vitor s ase kina then s pabe then s er slice header theke pointer er maddhome array er adress e pochabe ja adress er first element er adress then adress er shate 2 jog korbe then j adress pabe tar value 5 kore dibe. akhon jokhon s[4]=5 ai line asbe asbe ar same process kore jog korar por dekbe j joto number cell pawa jacche ta array er moddhe pore nah tokhon run time error dibe karon array er length holo 3 but amra access korte cacchi 4 ja length er bairer cell. dhori araay te ase adress ad18->1,ad19->2 ,ad20->3 kintu jokhon s er pointer ke follow korbe o khuje pabe ad18 tar sate 4 jog korle hobe ad22 ar ad22 name cell ta length er awotavukto nah.
// ---------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var s []int  // empty or nil slice
// fmt.print(s) 
// }

// simulation: main call hobe akta sell er nam rakhbe s tate ptr-nil,cap-0,len-0 rakhbe then jokhon fmt.print(s) aii line e asbe tokhon dekbe s ase kina s pabe er moddhe ptr pabe nill so [] print korbe

// ----------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var x []int          // [] , len = 0, cap = 0
// 	x = append(x, 1)     // [1], len = 1, cap = 1
// 	x = append(x, 2)     // [1, 2], len = 2, cap = 2
// 	x = append(x, 3)     // [1, 2, 3]
    
// 	y := x               // y and x share the same underlying array

// 	x = append(x, 4)     // may reallocate depending on capacity
// 	y = append(y, 5)     // append via y

// 	x[0] = 10             // modify underlying array

// 	fmt.Println(x)        // [10 2 3 5]
// 	fmt.Println(y)        // [10 2 3 5]
// }

// -> slice underlying array rules : 1024 porjonto 100% kore incrise then 25%
// task simulation
// -----------------------------------------------------------------------------
// package main

// import "fmt"

// func changeSlice(p []int) []int {
// 	p[0] = 10 // hr-24 len-3 cap-6 [10,6,7]
// 	p = append(p, 11) // hr-24 len-4 cap-6 [10,6,7,11]
// 	return p
// }

// func main() {
// 	x := []int{1, 2, 3, 4, 5} //r-4
// 	x = append(x, 6) // hr-20 len-6 cap-10 [1, 2, 3, 4, 5,6]
// 	x = append(x, 7) // hr-20 len-7 cap-10 [1, 2, 3, 4, 5,6,7]

// 	a := x[4:] // hr-24 len-3 cap-6 [5,6,7]

// 	y := changeSlice(a) // hr-24 len-4 cap-6 [10,6,7,11]

// 	fmt.Println(x) //[1, 2, 3, 4, 10, 6, 7]
// 	fmt.Println(y) // [10,6,7,11]
// 	fmt.Println(a) // [10,6,7]
// 	fmt.Println(x[0:8]) // [1, 2, 3, 4, 10, 6, 7, 11]
// }

// simulation
// -------------------------------------------------------------------------------
// package main

// import "fmt"

// // variadic function
// func print(numbers ...int) {
// 	fmt.Println(numbers)
// 	fmt.Println(len(numbers))
// 	fmt.Println(cap(numbers))
// }

// func add(a int, b int) {

// }

// func main() {
// 	print(5, 6, 7, 8, 9)
// }
