package main

import "fmt"

func main(){
    fmt.Println("Введите массу в килограммах: ")
    var num float64
    fmt.Scanf("%f", &num)

    answer := num / 1000

    fmt.Println(answer,"тонн")    
}
