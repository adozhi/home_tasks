package main

import "fmt"

func main(){
	var k , day int
	fmt.Scan(&k)
	fmt.Scan(&day)
	fmt.Println(k % 7 + day - 1)
}

