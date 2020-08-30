/*

Insertion Sort, ou ordenação por inserção, é um algoritmo de ordenação que, dado uma estrutura (array, lista) constrói uma matriz final com um
elemento de cada vez, uma inserção por vez.Assim como algoritmos de ordenação quadrática, é bastante eficiente para problemas com pequenas entradas,
sendo o mais eficiente entre os algoritmos desta ordem de classificação.

Melhor caso: O(n), quando a matriz está ordenado;
Médio caso: O(n²/4), quando a matriz tem valores aleatórios sem ordem de classificação (crescente ou decrescente);
Pior caso: O(n²), quando a matriz está em ordem inversa, daquela que deseja ordenar.

*/
package main
 
import (
    "time"
)
  
func insertion_sort(items []int) ([]int, int, float64) {
	var count int
	var inicio int64	
	var final int64
	var tempofinal float64

	var n = len(items)
	
	inicio = time.Now().UnixNano();

    for i := 1; i < n; i++ {
		count++;
		j := i
		count++;
        for j > 0 {
			count++;
            if items[j-1] > items[j] {
				count++;
				items[j-1], items[j] = items[j], items[j-1]
				count++;
            }
			j = j - 1
			count++;
        }
	}
	
	tempofinal = float64(final - inicio)/1000000;

	return items,count,tempofinal;
}