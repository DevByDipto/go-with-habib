package main

import "fmt"

func main(){
	var a = 10
	var b int = -20  
	var c int8 = 30  
	// 1. question aikhane uporer ai 3 ta variable er type er moddhe parthokko kii ?

	/*ans: ai 3 tar moddhe parthokko holo j a=10 mane a akta int type data ja ram e full jaiga nibe artahth amar computer jodi hoi 32 bit er tar mane ram er akta cell e 32 ta binary digite rakha jabe so jokhon ami a=10 likhbo tokhon ram ke 8 8 8 8 4 vage 8 kore vag kore tate 10 er binary man ta rakhbe ar khali bit gulo ke 0 diye fill up kore dibe. same for var b int = -20 just aikhane negative value rakhbe.ar var c int8 = 30 er mane holo amar ram joto bit er e hoknah keno c shudu akti cell er 8 bit er jaiga nibe.jodi dhori amar ram ta 32 bit ar tahole se 8 bit er nam rakhbe c baki 24 bit khali rakhbe.
	*/ 
	// 2. jodi amar cmputer 32 bit er hoi tahole seta ram e variable assign korar somoi kii provap felbe?
	/*ans: 1 number question e answer dewa ase.
	*/ 

	var d unit = 28 // unsigned 
	// 3. question aikhane uporer ai 4 ta variable er type er moddhe parthokko kii ?
	/*ans: unit type er mane holo aikhane unsigned value rakha jabe mane sudu positive value rakha jabe.
	*/ 
	var e float64 = 54.644
    // 4. question aikhane uporer ai 5 ta variable er type er moddhe parthokko kii ?
	/*ans: float64 ar mane holo aikhane just flotation value rakha jabe.
	*/  
	// 5. amar computer jodi 32 bit er hoi but ami data type bosaisi 64 bit er tahole seta memory te kivabe effect felbe ?
	/*ans: ai condition e go run time duita cell ja 32 bit sai 2 ta ke milai 1 ta cell banabe ar tar nam dibe a ar tate value rekhe dibe.
	*/ 

	// 6. init8 mane aikhane -128 theke 127 porjonto value rakha jabe ar unit8 mane aikhane 255 porjonto rakha jabe aikhane 256 er binary to 8 bit tahole aita init8 e rakha jabe nah keno ?
	// ans: chatgpt

	var flag bool = true
	// 7. bool koto bit jaiga dokhol kore ?
	// ans: 8bit
	// 8. formater ki keno use kori ?
	// ans: memory te data thake binary formate e akhon ai binary formate data ke amra jodi tader original formate e dekte cai tar jonno amra formater use kori
	// 9. ami jodi akti variable flotation e likhi but use kori %d tahole amake desimal value dekhabe ? ans : chatgpt
}