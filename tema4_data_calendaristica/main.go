package main

import "fmt"
import c "./calendar"

func main() {

	data1 := c.DataCalendaristica{Day: 30, Month: 07, Year: 2021}
	data2 := c.DataCalendaristica{Day: 14, Month: 04, Year: 2021}
	c.ShowDate(data1)
	c.ShowDate(data2)


	fmt.Println("----------")
	fmt.Println("Original:")
	c.ShowDate( data1)
	fmt.Println("----------")
	fmt.Println("Minus 1 day:")
	c.ShowDate(data1.MinusOneDay())
	fmt.Println("Minus 1 month:")
	c.ShowDate(data1.MinusOneMonth())
	fmt.Println("Minus 1 year:")
	c.ShowDate(data1.MinusOneYear())
	fmt.Println("----------")
	fmt.Println("Original:")
	c.ShowDate(data1)
	fmt.Println("----------")
	fmt.Println("Plus 1 day:")
	c.ShowDate(data1.PlusOneDay())
	fmt.Println("Plus 1 month:")
	c.ShowDate(data1.PlusOneMonth())
	fmt.Println("Plus 1 year:")
	c.ShowDate(data1.PlusOneYear())



	fmt.Println("Numar de zile intre",data1,data2)
	fmt.Println(c.NumberOfDaysBetweenTwoDates(data2,data1))

}
