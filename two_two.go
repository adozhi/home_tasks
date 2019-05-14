package main

import "fmt"

func main(){
    fmt.Println("Введите массу в килограммах: ")
    var num float64
    fmt.Scanf("%f", &num)

    answer := num / 100

    fmt.Println(answer,"полных центнеров ")    
}
