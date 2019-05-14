package main

import "fmt"

func main(){
    fmt.Println("Введите количество секунд: ")
    var num int
    fmt.Scan(&num)

	answer_one := (num / 60) / 60
	answer_two := (num - 3600 * (num/3600)) / 60
	answer_three := num % 60
	fmt.Println(answer_one,"полных часов с начала суток") 
	fmt.Println(answer_two,"полных минут с начала очередного часа") 
	fmt.Println(answer_three,"полных секунд с начала очередной минуты")    
}

