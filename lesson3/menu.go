package main

import "fmt"

func main() {
	//vars
	menu := [3]string{"Маргарита", "Дьябола", "Гавайская"}
	size := [2]string{"10 дюймов", "15 дюймов"}
	var client_choice int
	var client_choice2 int
	var name_pizza []string
	var sum int = 0
	var continue_cycle int
	var end int
	//
	fmt.Println("Добро пожаловать в пиццерию")

	//infinity cycle

	for {

		fmt.Println("Меню: \nдля выбора наберите цифру")
		for i := 0; i < len(menu); i++ {
			fmt.Println(i+1, " - ", menu[i])
		}

		fmt.Scan(&client_choice)

		//correct choice

		if client_choice > 0 && client_choice < 4 {
			fmt.Println("Выберите размер пиццы\nдля выбора наберите цифру")
			for i := 0; i < len(size); i++ {
				fmt.Println(i+1, " - ", size[i])
			}
			fmt.Scan(&client_choice2)
			continue_cycle = 0

			//conditions with a choice

			switch client_choice {
			case 1:
				if client_choice2 == 1 {
					name_pizza = append(name_pizza, "Маргарита/10дюймов")
					sum += 200
				} else if client_choice2 == 2 {
					name_pizza = append(name_pizza, "Маргарита/15дюймов")
					sum += 270
				} else {
					fmt.Println("Error, write retry")
					continue_cycle = 1
				}
			case 2:
				if client_choice2 == 1 {
					name_pizza = append(name_pizza, "Дьябола/10дюймов")
					sum += 225
				} else if client_choice2 == 2 {
					name_pizza = append(name_pizza, "Дьябола/15дюймов")
					sum += 280
				} else {
					fmt.Println("Error, write retry")
					continue_cycle = 1
				}
			case 3:
				if client_choice2 == 1 {
					name_pizza = append(name_pizza, "Гавайская/10дюймов")
					sum += 270
				} else if client_choice2 == 2 {
					name_pizza = append(name_pizza, "Гавайская/15дюймов")
					sum += 350
				} else {
					fmt.Println("Error, write retry")
					continue_cycle = 1
				}
			default:
				fmt.Println("Error, try retry")
			}
		}
		if continue_cycle != 1 {
			fmt.Println("Продолжить выбор? Да - 1, Нет - 0")
			fmt.Scan(&end)
			if end == 0 {
				break
			}
		}
	}

	//check list
	fmt.Println("Название:")
	for i := 0; i < len(name_pizza); i++ {
		fmt.Println(name_pizza[i])
	}
	fmt.Println("Cумма: ", sum)

}
