package main

import "fmt"

import "time"

func main() {
	date := time.Now()
	fmt.Println("Welcome to out pay terminal")
	fmt.Println("Choose your operator")
	var mobile [4]string = [4]string{"Megafon", "Beeline", "Babilon-M", "Tcell"}
	for i := 0; i < 4; i++ {
		fmt.Println(string(i+1) + " - " + mobile[i])
	}
	var sum int
	var operators int
	var number int
	fmt.Scan(&operators)
	///Megafon
	if operators == 1 {
		fmt.Println("Enter your number Megafon:")
		fmt.Scan(&number)
		if number/10000000 == 88 || number/10000000 == 90 {
			fmt.Println("Enter sum:")
			fmt.Scan(&sum)
			if sum > 0 {
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
				fmt.Println()
				fmt.Println("Megafon")
				fmt.Println("number ", number)
				fmt.Println("sum ", sum)
				fmt.Println("Date :", date.Format("01/02/2006"))
				fmt.Println("OperationStatus : Successful")

				fmt.Println()
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
			}

		}

	}
	///Beeline
	if operators == 2 {
		fmt.Println("Enter your number Beeline:")
		fmt.Scan(&number)
		if number/10000000 == 88 || number/10000000 == 90 {
			fmt.Println("Enter sum:")
			fmt.Scan(&sum)
			if sum > 0 {
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
				fmt.Println()
				fmt.Println("Beeline")
				fmt.Println("number ", number)
				fmt.Println("sum ", sum)
				fmt.Println("Date :", date.Format("01/02/2006"))
				fmt.Println("OperationStatus : Successful")

				fmt.Println()
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
			}

		}

	}
	///Babilon-M
	if operators == 3 {
		fmt.Println("Enter your number Babilon-M:")
		fmt.Scan(&number)
		if number/10000000 == 91 {
			fmt.Println("Enter sum:")
			fmt.Scan(&sum)
			if sum > 0 {
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
				fmt.Println()
				fmt.Println("Babilon-M")
				fmt.Println("number ", number)
				fmt.Println("sum ", sum)
				fmt.Println("Date :", date.Format("01/02/2006"))
				fmt.Println("OperationStatus : Successful")

				fmt.Println()
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
			}

		}

	}
	///Tcell
	if operators == 4 {
		fmt.Println("Enter your number Tcell:")
		fmt.Scan(&number)
		if number/10000000 == 93 {
			fmt.Println("Enter sum:")
			fmt.Scan(&sum)
			if sum > 0 {
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
				fmt.Println()
				fmt.Println("Tcell")
				fmt.Println("number ", number)
				fmt.Println("sum ", sum)
				fmt.Println("Date :", date.Format("01/02/2006"))
				fmt.Println("OperationStatus : Successful")

				fmt.Println()
				for i := 0; i < 20; i++ {
					fmt.Print("-")
				}
			}

		}

	}
}
