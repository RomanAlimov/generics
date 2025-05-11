package main

import "fmt"

// Дженерики(generics) один пример:

func main() {
	s := []float64{0.6, 1.0, 44.2}
	g := []int64{1, 5, 7, 3, 1, 46, 7, 2, 131, 414}

	fmt.Println(sum(s)) // 45.800000000000004
	fmt.Println(sum(g)) // 617
}

func sum[v int64 | /* это оператор union*/ float64](num []v) v { // констреинт с доступными типами
	var sum v                 // сумма для v
	for _, num := range num { // перебор по num
		sum += num // прибавлявем сумму к цифре
	}
	return sum // возвращаем сумму для функции main()
}
