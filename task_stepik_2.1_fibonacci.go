/* Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn. */

package main
import "fmt"

func fibonacci(n int) int {
	var f []int
	var num int
	num = 1
	f = append(f, 0, 1)
	for i := 2; i <= n; i++ {
		num = f[i-1] + f[i-2]
		f = append(f, num)
	}
	return num
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(fibonacci(num))
}
