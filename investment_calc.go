package main

import (
	"fmt"
)



func main()  {
 var revenue float64
 var expenses float64
 var taxRate float64


   revenue = outPutText("revenue: ")

    expenses = outPutText("expenses: ")

    taxRate = outPutText("taxRate: ")

    calculateEbt(revenue, taxRate, expenses)


   
 }
 

func outPutText(text string) float64 {
    var userInput float64
   fmt.Print("Enter: ",text)
    fmt.Scanln(&userInput)

    return userInput
} 

func calculateEbt(revenue ,taxRate , expenses float64) (ebt ,profit ,ratio float64) {
    ebt = revenue - expenses
    profit = ebt * (1 - taxRate/100)
    ratio = ebt/profit

    fmt.Println("Earnings before tax: ", ebt)
    fmt.Println("Profit: ", profit)
    fmt.Println("Ratio: ", ratio)

    return ebt, profit, ratio
}