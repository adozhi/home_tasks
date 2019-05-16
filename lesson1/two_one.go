package main

import "fmt"

func main(){
    fmt.Println("Введите число в метрах: ")
    var num float64
    fmt.Scanf("%f", &num)

    answer := num / 100

    fmt.Println("Число полных метров в нем ",answer)    
}