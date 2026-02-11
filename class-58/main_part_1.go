package main

import (
	"fmt"
	"os"
)

// --- Interfaces ---

type People interface {
	PrintDetails()
	ReceiveMoney(amount float64) float64
}

type BankUser interface {
	WithdrawMoney(amount float64) float64
}

// --- Struct and Methods ---

type user struct {
	Name  string
	Age   int
	Money float64
}

// Receiver functions (methods) for the user struct
func (obj user) PrintDetails() {
	fmt.Printf("Name: %s, Age: %d, Money: %.2f\n", obj.Name, obj.Age, obj.Money)
}

func (obj user) ReceiveMoney(amount float64) float64 {
	obj.Money = obj.Money + amount
	return obj.Money
}

func (obj user) WithdrawMoney(amount float64) float64 {
	obj.Money = obj.Money - amount
	return obj.Money
}

// --- Main Execution ---

func main() {
	// usr1 assigned to People interface
	var usr1 People
	usr1 = user{
		Name:  "Habibur Rahman",
		Age:   30,
		Money: 10.08,
	}
// task1: aikhane user1 ke user type people type duitai bola jabe.
/*
ans 1: aikhane user people ke extand korse ba dashotto shikar korse tai user1 ke user type ba people type duitai bola jai jmn afrika jodi france er bashida hote cai tahole oder france goverment er dashotto shikar korte hobe france er niyom mante hobe tokhon afrikanra nijeder afrikan abar frances o bolte parbe tmn e aikhane.
*/
// task 2: aikhane usr1.PrintDetails(),usr1.ReceiveMoney(100) call kora jabe but usr1.WithdrawMoney(10) call kora jabe nah keno ta bujo
/*
ans2 : aikhane user struct er method PrintDetails,ReceiveMoney,WithdrawMoney 3 tai kintu user theke creat user1 se people interface ke extand korse baa dashotto sikar korse.jmn afrika jodi france er bashida hote cai tahole oder france goverment er dashotto shikar korte hobe france er niyom mante hobe tokhon afrikanra nijeder afrikan niyom o mante parbe abar france er niyom o mante hobe but upore thakbe france karon france er dashotto shikar korse. j j niyom france ar afrikan niyom er moddhe dondo toiri korbe se se niyomer khetre france er neyomke pradhanno dite hobe. tai aikhane people boltese or method holo sudu 2 ta PrintDetails,ReceiveMoney ar user struct boltese or method 3 ta kintu jehuru people mane france boltese 2 ta so oke akhon vapte hobe or o mane user1 er method 2 ta
*/
	// usr2 using short declaration
	usr2 := user{
		Name:  "Kalilur Rahman",
		Age:   30,
		Money: 10.08,
	}

	// usr3 assigned to BankUser interface
	var usr3 BankUser
	usr3 = user{
		Name:  "rahimullah",
		Age:   100,
		Money: 100000.24,
	}

	// Using BankUser method
	usr3.WithdrawMoney(10)
// task 3: aikhane usr3.WithdrawMoney(10) call kora jabe but usr3.PrintDetails(),usr3.ReceiveMoney(100) call kora jabe nah keno bujo.
	// --- Type Assertion ---
	// Converting the interface 'usr3' back to the concrete 'user' struct
	obj, ok := usr3.(user)// aikhane check kortesi user3 ki vitore vitore user struct kina
	if !ok {
		fmt.Println("Sorry usr3 is not type of user struct")
		os.Exit(1)
	}

	// Now we can access methods from both interfaces via 'obj'
	// task 4: aikhane keno object 3 ta method ke call korte partese?
	/*
	ans 4: obj mane aikhane user3 j kina BankUser er dashotto mene nisilo se abar nijer sadhi notai fire gese ar user struct er method chilo 3 ta tai aikhane obj 3 ta method ke e call korte parche
	*/
	obj.PrintDetails()
	obj.ReceiveMoney(10)
	obj.WithdrawMoney(100)

	// Final Printouts
	usr1.PrintDetails()
	usr2.PrintDetails()
	usr1.ReceiveMoney(100)
}