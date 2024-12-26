package main

import (
	"math"
	"fmt"
)



func main()  {
 var investmentAmount = 1000
 var annualInterestRate = 4.25
 var years = 1

 var futureInvestmentValue = float64(investmentAmount) * math.Pow(1 + annualInterestRate / 100,float64(years))
 
 fmt.Println(futureInvestmentValue)
 }
 

