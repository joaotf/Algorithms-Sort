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

func findMinAndMax(a []int) (max int) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func Counting_Sort(A []int, k int) (int, float64, []int) {
	var count int
	inicio := time.Now().UnixNano() / int64(time.Millisecond)

	C := make([]int, k+1)
	B := make([]int, len(A))

	for j := 0; j < len(A); j++ {
		C[A[j]]++
		count++
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1]
		count++
	}

	for j := len(A) - 1; j >= 0; j-- {
		n := A[j]
		B[C[n]-1] = n
		C[n]--
		count++
	}
	final := time.Now().UnixNano() / int64(time.Millisecond)

	tempo := float64(final - inicio)
	return count, tempo, B
}
