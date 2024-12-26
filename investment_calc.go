package main

import (
	"fmt"
	"math"
)



func main()  {
 var investmentAmount, years float64 = 1000,10
  annualInterestRate := 4.25


 futureInvestmentValue := investmentAmount * math.Pow(1 + annualInterestRate / 100,years)

 fmt.Println(futureInvestmentValue)
 }
 

