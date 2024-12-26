package main

import (
	"fmt"
)



func main()  {
 var investmentAmount, actualReturn float64 = 1000,10

  fmt.Print("Enter Actual return amount: ")
  fmt.Scan(&actualReturn)
  
  fmt.Print("Enter investment amount: ")
  fmt.Scan(&investmentAmount)


 

 profits := actualReturn - investmentAmount

 if profits > 0 {
 fmt.Print("You made a profit of: ")
 fmt.Println( profits)

 } else {
	 fmt.Print("You made a loss of: ")
	 fmt.Println( profits)
 }

 }
 

