package main

import "fmt"


func main(){
	age := 32 

	agePointer := &age
	fmt.Println("Age pointer:", agePointer)
	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println("Adult years:", adultYears)
}


func getAdultYears(age int) int{
   return age - 18
}