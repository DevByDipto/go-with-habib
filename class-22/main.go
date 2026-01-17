// package main 
// import "fmt"


// func main(){

// 	add := func(a int, b int){
//       c := a + b
// 				fmt.Println(c)
// 	}

// 	add(2,3)
// }

// func init(){
// 	fmt.Println("i well be call first")

// }

// // task : ram e code ta kivabe exicute hobe ?
//----------------------------------------------------------------
// package main 
// import "fmt"

// func sum(){
// 	add(2,4)
// }

// func add(a int, b int){
// 	fmt.Println(a+b)

// }

// func main(){
// sum()
// 	add := func(a int, b int){
//       c := a + b
// 				fmt.Println(c)
// 	}

// 	add(2,3)
// }

// func init(){
// 	fmt.Println("i well be call first")
// }
// // task : ram e code ta kivabe exicute hobe ?
//----------------------------------------------------------------

package main 
import "fmt"


func main(){

	add(2,3)

	add := func(a int, b int){
      c := a + b
				fmt.Println(c)
	}

	add(2,3)
}

func init(){
	fmt.Println("i well be call first")

}

// // task : ram e code ta kivabe exicute hobe ?

// aita error dibe karon local scope e function variable diclaer korar age access kora jai nah.