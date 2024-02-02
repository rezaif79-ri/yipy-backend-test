package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var arr = make([]int, 0)
	fmt.Print("Input array: ")
	if scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		for i := range input {
			v, err := strconv.Atoi(input[i])
			if err != nil {
				fmt.Println("error parsing input: ", err)
				return
			}
			arr = append(arr, v)
		}
	}

	var maxProfit int = GetMaxProfit(arr)
	fmt.Println(maxProfit)
}

// GetMaxProfit takes slice integer to returning max profit (max value between)
func GetMaxProfit(stocks []int) int {
	var profit int = 0
	var min = math.MaxInt

	for i := range stocks {
		if min > stocks[i] {
			min = stocks[i]
		}

		current := stocks[i] - min
		if current > profit {
			profit = current
		}
	}
	return profit
}
