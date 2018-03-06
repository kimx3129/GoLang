package main

import "fmt"
//
//func MakeGreeter() func() string{
//	return func() string{
//		return "Hello World!"
//	}
//}

//func wrapper() func() int{
//	var x int
//	return func() int{
//		x++
//		return x
//	}
//}
//
//func filter(numbers []int, callback func(int) bool) []int{
//	xs := []int{}
//	for _, n := range numbers{
//		if callback(n){
//			xs = append(xs, n)
//		}
//	}
//	return xs
//}

func main(){
	//data := []float64{43, 56, 76, 112, 45, 57}
	//n := average(data...)
	//fmt.Println(n)

	//haha := MakeGreeter()
	//fmt.Println(haha())

	//increment := wrapper()
	//fmt.Println(increment())
	//fmt.Println(increment())

	//xs := filter([]int{7,3,8,3,5,8,4,1,7,9,0}, func(n int) bool{
	//	return n > 3
	//})
	//
	//fmt.Println(xs)

	// Pass by Value
	//age := 44
	//fmt.Println(&age)
	//
	//ChangeMe(&age)
	//
	//fmt.Println(&age)
	//fmt.Println(age)

	m := make([]string, 1, 25)
	fmt.Println(m)
	ChangeMe(m)
	fmt.Println(m)
}
//
//func average(sf ...float64) float64{
//	total := 0.0
//	for _, v := range sf{
//		total += v
//	}
//	return total / float64(len(sf))
//}

//func ChangeMe(z *int){
//	fmt.Println(z)
//	fmt.Println(*z)
//
//	*z = 24
//	fmt.Println(z)
//	fmt.Println(*z)
//}

func ChangeMe(z []string){
	z[0] = "Todd"
	//z[1] = "Sungmin"
	fmt.Print(z)
}