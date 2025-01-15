package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

type radix struct {
	base      int
	name      string
	shorthand string
}

var radii = []radix{
	{2, "binary", "0b"},
	{8, "octal", "0o"},
	{10, "decimal", ""},
	{16, "hexadecimal", "0x"},
}

const MinValue int = 0
const MaxValue int = 64

func main() {
	var correct int = 0
	var attempted int = 0
	for {
		// generate the random radii
		var fromRadix, toRadix radix = select2(radii)

		// generate the random value
		var value int64 = int64(MinValue + rand.IntN(MaxValue-MinValue+1))

		// get its values as strings
		var fromValue string = strconv.FormatInt(value, fromRadix.base)
		var toValue string = strconv.FormatInt(value, toRadix.base)

		// print the question to the user
		fmt.Printf("%s%s -> %s", fromRadix.shorthand, fromValue, toRadix.shorthand)

		// get the input from the user
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Could not read input, please try again.")
			continue
		}

		// if the user said quit
		if input == "q" || input == "Q" {
			break
		}

		// if the user guessed
		attempted++
		if input == toValue {
			correct++
			fmt.Printf("↳ Correct ✅ - %d/%d (%.2f%%)\n",
				correct, attempted, float32(correct)/float32(attempted)*100,
			)
		} else {
			fmt.Printf("↳ Incorrect ❌ | The correct answer was %s%s",
				toRadix.shorthand,
				toValue,
			)

			if toRadix.base != 10 {
				fmt.Printf(" (or %d in decimal)", value)
			}

			fmt.Printf(" - %d/%d (%.2f%%)\n",
				correct, attempted, float32(correct)/float32(attempted)*100,
			)
		}
	}
}

func select2[T any](s []T) (T, T) {
	var maxSize int = len(s)

	// check if there are less than 2
	if maxSize < 2 {
		panic("slice isn't long enough!")
	}

	// generate the two random indices
	var index1 int = rand.IntN(maxSize)
	var index2 int = rand.IntN(maxSize)

	// try again if collision
	for index1 == index2 {
		index2 = rand.IntN(maxSize)
	}

	return s[index1], s[index2]
}
