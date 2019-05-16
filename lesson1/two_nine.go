package main

import "fmt"

func main(){
	var n, x int
	fmt.Print("Введите месяц:")
	fmt.Scan(&n)
	x = (n % 12) + 1
	fmt.Println("Месяц: ", x)
}

