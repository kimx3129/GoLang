package main

import "fmt"

//func question_one(val int) (int, bool, bool){
//	return val / 2, val % 2 == 0, val % 3 == 0 && val % 5 == 0
//}

//func question_three(num ...int) int{
//	greatest := 0
//	for _, val := range num{
//		if val > greatest{
//			greatest = val
//		}
//	}
//	return greatest
//}

func foo(num ...int){
	fmt.Printf("Type : %T\n", num)
	fmt.Println(num)
}

func main(){

	//question_one_func_exp := func(val int) (int, bool){
	//	return val / 2, val % 2 == 0
	//}
	//a, b := question_one_func_exp(12)
	//fmt.Println(a,b)

	//ls := []int{1,2,3,4,5,8,11,13,16,18,44,8,3,1,43,3,74,34,23}
	//a := question_three(ls...)
	//fmt.Println(a)

	foo(1,2)
	foo(1,2,3)
	aSlice := []int{1,2,3,4}
	foo(aSlice...)
	foo()
}
