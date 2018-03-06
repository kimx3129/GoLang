package main

import "fmt"

func main(){
	//m := make(map[string]int)
	//m["name"] = 23
	//m["Age"] = 55
	//
	//delete(m, "Age")
	//fmt.Println(m)
	//
	//a, _ := m["name"]
	//fmt.Println(a)
	//
	//n := map[int]string{33:"Kim", 31:"Grace"}
	//fmt.Println(n)
	//myGreeting := map[string]string{
	//	"Sungmin":"Hello",
	//	"Grace":"Jesus",
	//}
	//myGreeting["Grace"] = "HHHHH"
	//fmt.Println(myGreeting)

	//myGreeting := map[int]string{
	//	0: "Good Morning",
	//	1: "Hello there",
	//	2: "Welcome",
	//}
	//
	//fmt.Println(myGreeting)
	//
	//if val, exists := myGreeting[5]; exists{
	//	delete(myGreeting, 5)
	//	fmt.Println(val)
	//	fmt.Println(exists)
	//
	//}

	// range loop
	carviEmployee := map[string]int{
		"Patrick":44,
		"Bradon":47,
		"Sungmin":43,
	}

	for k, v := range carviEmployee{
		fmt.Println(k, '-', v)
	}

}
