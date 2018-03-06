package main

import (
	"fmt"
	"io"
	"crypto/rand"
)


func main(){
	//mySlice := []int{1,2,3,4,5}
	//var x[42] int
	//
	//fmt.Printf("%T\n", mySlice)
	//fmt.Printf("%T\n", x)
	//fmt.Println(mySlice)
	//
	//fmt.Println("Sungmin"[0])

	//mySlice := make([]int, 0, 5)
	//
	//fmt.Println("---------------")
	//fmt.Println(mySlice)
	//fmt.Println(len(mySlice))
	//fmt.Println(cap(mySlice))
	//fmt.Println("---------------")
	//
	//for i := 0 ; i < 40 ; i++{
	//	mySlice = append(mySlice, i)
	//	fmt.Println("Len : ", len(mySlice), "cap : ", cap(mySlice), "value : ", mySlice[i])
	//}

	//Myslice := []int{1,2,3,4,5}
	//Myslice2 := []int{6,7,8,9,10}
	//
	//Myslice = append(Myslice, Myslice2...)
	//
	//fmt.Println(Myslice)

	//records := make([][]int, 0, 5)
	//
	//for i := 0 ; i < 3 ; i++{
	//	transaction := make([]int, 0)
	//	for j := 0 ; j < 6 ; j++{
	//		transaction = append(transaction, j)
	//	}
	//	records = append(records, transaction)
	//}
	//
	//fmt.Println(records)

	//short-hand way
	//student := []string{}
	//student_two := make([]string, 2)
	////students := [][]string{}
	//
	//student_two = append(student_two, "Sungmin")
	//student_two = append(student_two, "grace")
	//student_two = append(student_two, "chris")
	//
	////use var
	//var student []string
	//var students [][]string
	//
	////use make
	//student := make([]string, 35)
	//students := make([][]string, 32)
	//fmt.Println(student_two)
	//fmt.Println(cap(student_two))

	//mySlice := make([]int, 1)
	//fmt.Println(mySlice)
	//mySlice[0] = 7
	//mySlice = append(mySlice, 111)
	//fmt.Println(mySlice)
	//mySlice[0]++
	//mySlice[1]++
	//fmt.Println(mySlice)

	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)
}