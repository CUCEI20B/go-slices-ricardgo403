package main

import "fmt"

func main() {
	var tamanio uint
	fmt.Scan(&tamanio)
	var mySlice = make([]uint, 0, tamanio)
	var sumatoria uint
	for i := 0; i < int(tamanio); i++ {
		var captured uint
		if _, err := fmt.Scan(&captured); err != nil {
			fmt.Println(err)
			return
		}
		mySlice = append(mySlice, captured)
	}
	for _, v := range mySlice {
		sumatoria += v
	}
	fmt.Println(sumatoria)
}
