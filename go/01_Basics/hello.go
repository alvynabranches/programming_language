// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	// var score int = 4
	// fmt.Println(score)
	// gives error if the variable is declared but not used
	// const score int = 400
	// fmt.Println(score)
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	thorRating := 3.
	fmt.Printf("%v, %T", thorRating, thorRating)

	var Ironman, CatAmerica string = "I am Ironman", "I am Captain America"
	fmt.Println(Ironman)
	fmt.Println(CatAmerica)

	var defaultValue int
	fmt.Println(defaultValue)

	var (
		spiderman = "I am spiderman"
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "I am antman"
	)
	fmt.Println(spiderman)
	fmt.Println(age)
	fmt.Println(powers)
	fmt.Println(antman)

}
