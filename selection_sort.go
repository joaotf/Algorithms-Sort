/*
A ordenação por seleção (do inglês, selection sort) é um algoritmo de ordenação baseado em se 
passar sempre o menor valor do vetor para a primeira posição (ou o maior dependendo da ordem requerida)
, depois o de segundo menor valor para a segunda posição, e assim é feito sucessivamente 
com os n − 1 elementos restantes, até os últimos dois elementos.
É composto por dois laços, um laço externo e outro interno. O laço externo serve para controlar o índice inicial e o interno percorre todo o vetor. Na primeira iteração do laço externo o índice começa de 0 e cada iteração ele soma uma unidade até o final do vetor e o laço mais interno percorre o vetor começando desse índice externo + 1 até o final do vetor.

*/

package main

import (
  "time"
)



func Selection_Sort(items []int) ([]int, int, float64) {
	var count int
	var inicio int64	
	var final int64
	var tempofinal float64

	var n = len(items)

	inicio = time.Now().UnixNano()
	for i := 0; i < n; i++ {
		count++
		var minIdx = i
		count++
		for j := i; j < n; j++ {
			count++
			if items[j] < items[minIdx] {
				count++
				minIdx = j
				count++
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
		count++
	}
	final = time.Now().UnixNano() 

	
  	tempofinal = float64(final - inicio)/1000000

	return items, count, tempofinal
}
