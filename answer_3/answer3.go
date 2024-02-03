package main

import "fmt"

// Question:

/*
	Untuk memenangkan hadiah untuk penjualan yang paling banyak terjual, teman saya
	dan saya akan menggabungkan “Cout Cookies” menjadi 1 pesanan. Setiap pesanan
	diwakili oleh “nomor pesanan” (bilangan bulat).

	Kami memiliki daftar pesanan yang sudah diurutkan secara numerik, dalam daftar.
	Buatlah fungsi untuk menggabungkan daftar pesanan kami menjadi 1 pesanan yang
	terurut. Tidak diperbolehkan menggunakan fungsi array yang telah ditentukan untuk
	menggabungkan dan mengurutkan (array_merge, array_sort, .sort(), .merge() atau
	fungsi serupa lainnya).
*/

func main() {
	// Solution No.1:
	// If the questions means to merge all the same "order name"
	// To 1 orders that contains all order id,
	// and the order id needs to be ordered from lowest to highest
	// The key considered to be order number
	var ordersIn map[int]string = map[int]string{
		1: "Cout Cookies", 2: "Noodles", 3: "Milks",
		4: "Fried Rice", 5: "Milks", 6: "Noodles",
	}

	var orders map[string][]int = make(map[string][]int)

	var sortSliceInt = func(num []int) []int {
		for i := 0; i < len(num)-1; i++ {
			for k := 1; k < len(num); k++ {
				if num[k] < num[i] {
					num[k], num[i] = num[i], num[k]
				}
			}
		}
		return num
	}
	for i := range ordersIn {
		if v, ok := orders[ordersIn[i]]; !ok {
			v = make([]int, 0)
			v = append(v, i)
			orders[ordersIn[i]] = v
		} else {
			v = append(v, i)
			v = sortSliceInt(v)
			orders[ordersIn[i]] = v
		}
	}

	fmt.Println(orders)
}
