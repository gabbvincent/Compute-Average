//Name: Vincet G.
//Date: 4/13/2020
//Descritpion: Compute average

package main

import "fmt"

func average(a, b int)(string, int, string,int,string,int){
//declare a variabel for avg as (a+b)/2 to find the average
avg := (a+b)/2

return "with a score of",a,"and a possible score of",b,"Your average score is",avg
}


func main() {
//declare variable for a and b as integer
var a int
var b int
//ask user for score and scan for a
fmt.Println("Enter your score")
fmt.Scanln(&a)
//ask user for points possible and scan for b
fmt.Println("Enter in the amount of points possible")
fmt.Scanln(&b)
//cal average(a, b)
 fmt.Println(average(a, b))
}