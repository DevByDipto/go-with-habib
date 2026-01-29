 package main

import "fmt"

func main(){
	arr := [6]string {"this","is","a","go","interview","question"}
fmt.print(arr)

s := arr[1:4] // slice nijer moddhe 3 ta value rakhe 
// 1. pointer-> slice kora array er value gulor first value baa element er adress
// 2. length-> slice array er moddhe koita value baa element rakhlo jmn aikhane len=3
// 3. capacity-> jekhan theke clice kora shuru korse se value ta hoso oi original array te ar koita value baa element ase setar totall number jmn amader aikhetre cap=5

fmt.print(s)

}

// task simulate : fmt.print(arr) ai porjonto to parbo then main stack fram e s name akta sell nibe tate akta info rakhbe ptn,len,cap then jokhon s print hobe tokhon dekbe main er moddhe s ase kina dekbe ase then s er value dekbe ptn somthing ar len-3 so ptn holo arr er slice start value oi jaiga theke 3 ta value tule ane print korbe. 
// 