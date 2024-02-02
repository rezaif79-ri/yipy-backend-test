package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	snum := GenerateSNum(n)

	for i := 1; i <= n; i++ {
		fmt.Println(GetFormatedSNum(i, snum))
	}
}

// GenerateSNum return slice string of number in ordered postion
// Takes n argument as how much number count + 3 (spaces needed for ** symbol and last number)
func GenerateSNum(n int) []string {
	var res []string
	for i := 1; i <= n+3; i++ {
		res = append(res, fmt.Sprint(i))
	}
	return res
}

// GetFormatedSNum return string formatted with question 1 requirement format (n * * n)
func GetFormatedSNum(n int, snum []string) string {
	var res []string
	if len(snum) > n+2 {
		res = append(res, snum...)
		res[n] = "*"
		res[n+1] = "*"
		return strings.Join(res, "")
	}
	return ""
}
