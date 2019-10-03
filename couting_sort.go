/*
Counting sort é um algoritmo de ordenação estável cuja complexidade é O(n).
As chaves podem tomar valores entre 0 e M-1. Se existirem k0 chaves com valor 0,então ocupam as primeiras
k0 posições do vetor final: de 0 a k0-1.
A ideia básica do counting sort é determinar, para cada entrada x, o número de elementos menor que x.
Essa informação pode ser usada para colocar o elemento x diretamente em sua posição no array de saída.
Por exemplo, se há 17 elementos menor que x, então x pertence a posição 18. Esse esquema deve ser ligeiramente
modificado quando houver vários elementos com o mesmo valor, uma vez que nós não queremos colocar eles
na mesma posição.
Complexidade(Pior Caso) : O( n + k )
Complexidade(Caso Médio) : O( n + k )
Complexidade(Melhor Caso) : O( n + k )
*/

package main

import (
	"time"
)


func findMinAndMax(a []int) (int,int) {
	var count1 int
	var max int
	max = a[0]
	count1++
	for _, value := range a {
		count1++
		if value > max {
			count1++
			max = value
			count1++
		}
	}
	return max,count1
}

func Counting_Sort(A []int, k int) (int, float64, []int) {
	var count int
	var inicio int64	
	var final int64
	var tempofinal float64

	inicio = time.Now().UnixNano() 
	C := make([]int, k+1)
	B := make([]int, len(A))

	for j := 0; j < len(A); j++ {
		count++
		C[A[j]]++
		count++
	}

	for i := 1; i <= k; i++ {
		count++
		C[i] += C[i-1]
		count++
	}

	for j := len(A) - 1; j >= 0; j-- {
		count++
		n := A[j]
		count++
		B[C[n]-1] = n
		count++
		C[n]--
		count++
	}
	final = time.Now().UnixNano()

	tempofinal = float64(final - inicio)/1000000
	return count, tempofinal, B
}