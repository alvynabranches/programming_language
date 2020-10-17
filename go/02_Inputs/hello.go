package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var myString string
	// fmt.Scanln(&myString)
	var name string = "alvyn"
	var a, _ = fmt.Println(name)
	fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// myname, _ := reader.ReadString('\n')
	// fmt.Println(myname)

	fmt.Print("Enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumRating)
}
