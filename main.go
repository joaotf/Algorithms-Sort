package main

import (
	"fmt"
	"math/rand"
	"time"
)

func perguntaproDoidao() int {
	var tamanho int
	fmt.Print("\n\nTamanho da ARRAY : ")
	_, err := fmt.Scan(&tamanho)
	if err == nil {
		fmt.Println()
	}
	return tamanho
}

func arraySort(tamanho int) []int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	a := make([]int, tamanho)
	for i := 0; i < len(a); i++ {
		a[i] = r1.Intn(20000)
	}
	return a
}

func main() {
	var tamanho int
	var array []int
	var max int
	var menu int

	fmt.Print("Menu\n1)Counting Sort\n2)Selection Sort\n3)Bubble Sort\nOpção:")
	_, err := fmt.Scan(&menu)
	if err == nil {
	}

	switch menu {
	case 1:
		{
			tamanho = perguntaproDoidao()
			array = arraySort(tamanho)
			max = findMinAndMax(array)
			fmt.Println("Array não ORDENADO : ", array)
			count, tempofinal, resultado := Counting_Sort(array, max)

			fmt.Println("\n\nCounting Sort (Finished) : ", resultado, "\n\nContador : ", count-max, "\nTempo : ", tempofinal)
		}
	case 2:
		{
			tamanho = perguntaproDoidao()
			array = arraySort(tamanho)
			fmt.Println("Array não ORDENADO : ", array)
			arrays, count, tempofinal := Selection_Sort(array)
			fmt.Println("\n\nSelection Sort (Finished) : ", arrays, "\n\nContador : ", count, "\nTempo : ", tempofinal)
		}
	case 3:
		{
			tamanho = perguntaproDoidao()
			array = arraySort(tamanho)
			fmt.Println("Array não ORDENADO : ", array)
			arrayd, count, tempofinal := Bubble_Sort(array)
			fmt.Println("\n\nBubble Sort (Finished) : ", arrayd, "\n\nContador : ", count, "\nTempo : ", tempofinal)
		}

	}

}
