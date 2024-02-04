package main

import "fmt"

// Question:
/*
	Anda memiliki daftar bilangan bulat, dan untuk setiap indeks Anda ingin menemukan
	produk dari setiap bilangan bulat kecuali bilangan bulat pada indeks itu. Buatlah fungsi
	get_data() yang mengambil daftar bilangan bulat dan mengembalikan daftar produk.
*/

type Product struct {
	ProductName string
	Price       int
}

func main() {
	var data []Product = init_data(10)

	getData, _ := get_data(5, data)
	fmt.Println(getData)
}

func get_data(index int, data []Product) ([]Product, error) {
	var res []Product
	res = append(data[:index], data[index+1:]...)
	return res, nil
}

func init_data(count int) []Product {
	var res []Product
	for i := 1; i <= count; i++ {
		res = append(res, Product{
			ProductName: fmt.Sprintf("Product %v", i),
			Price:       1000 * i,
		})
	}
	return res
}
