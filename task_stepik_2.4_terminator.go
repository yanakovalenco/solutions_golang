/*Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.
Если значение On == false, то оба метода вернут false.
Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true), если его нет,
то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.
*/
package main

import "fmt"

type Mystruct struct {
	On    bool
	Ammo  int
	Power int
}

func (s *Mystruct) Shoot() bool {
	if s.On == false {
		return false
	}
	if s.Ammo != 0 {
		s.Ammo--
		return true
	} else {
		return false
	}
}

func (s *Mystruct) RideBike() bool {
	if s.On == false {
		return false
	}
	if s.Power != 0 {
		s.Power--
		return true
	} else {
		return false
	}
}

func main() {
	testStruct := &Mystruct{}
	if testStruct.Shoot() && testStruct.RideBike() {
		fmt.Println("Wrong, check again")
	} else {
		fmt.Println("Passed. I will be back!")
	}
}
