package main

import "fmt"

func main() {
	const initial = 100000.0 // начальный вклад
	const rate = 12.0        // годовая ставка, %
	const monthsInYear = 12  // количество месяцев в году

	monthly := rate / monthsInYear / 100 // ежемесячная доля: 12/12/100 = 0.01

	//баланс за каждый месяц
	balance := initial
	for month := 1; month <= monthsInYear; month++ {
		balance = balance * (1 + monthly)
		fmt.Printf("месяц %2d: %.2f\n", month, balance)
	}

	//через сколько месяцев удвоение
	balance = initial // сброс, иначе считаем не с нуля
	months := 0
	for balance < initial*2 { // for только с условием
		balance = balance * (1 + monthly)
		months++
	}

	fmt.Printf("Вклад удвоится через %d мес. Баланс: %.2f\n", months, balance)
}

//Вопрос: почему в реальных системах деньги не хранят во `float64`?
// float64 теряет точность при работе с десятичными дробями
