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

// func main(){     //    18     19   20  21       22         23
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

package main

import "fmt"

func main() {

    s1 := []int{1, 2, 3}  // its call sclice litaral

    fmt.Println(s1)
    fmt.Println(len(s1))
    fmt.Println(cap(s1))
}

// task simulation: main porjonto parbo then main call hobe stack fram toiri hobe tate 3 ta sell nibe se sell e 123 value bosbe then s1 name akta cell alocate hobe tate ptr,cap,len slice header rakhbe ar porer ta parbo.