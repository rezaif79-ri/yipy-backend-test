package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Question:
/*
	Pada sebuah penerbangan terdapat sistem menonton, terkadang sebelum film tersebut
	habis kereta telah berhenti. Buatlah fungsi yang menentukan hasil penjumlahan 2
	durasi film dari koleksi film yang hasil penjumlahan tersebut apakah sama dengan
	lamanya perjalanan n yang diinputkan.

	* Diasumsikan setiap orang menonton setidaknya 2 film dan filmnya tidak boleh sama
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Input n: ")
	var n int
	if scanner.Scan() {
		input, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("error parsing input: ", err)
			return
		}
		n = input
	}

	// films contains list of film, key(string) used as the title and value(int) as the duration
	var films map[string]int = map[string]int{
		"Harry Potter": 120,
		"Finding Nemo": 95,
		"John Wick":    110,
		"UP":           92,
	}
	var titles []string
	for i := range films {
		titles = append(titles, i)
	}

	var found bool
	for i := range titles {
		if i < len(titles)-1 {
			for x := i + 1; x < len(titles); x++ {
				if films[titles[i]]+films[titles[x]] == n {
					found = true
					fmt.Println(titles[i], "and", titles[x], "is", n, "minutes")
				}
			}
		}
	}
	if !found {
		fmt.Println("The flight ended before or after 2 films")
	}
}
