package main

import "fmt"

func main(){
    fmt.Println("Введите расстояние в метрах: ")
    var num float64
    fmt.Scanf("%f", &num)

    answer := num / 1000

    fmt.Println(answer,"километров")    
}
