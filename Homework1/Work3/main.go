package main

import "fmt"

func main() {
    const initial = 100000.0  // начальный вклад
    const rate = 12.0         // годовая ставка, %

    monthly := rate / 12 / 100  // ежемесячная доля: 12/12/100 = 0.01

    // ---------- Часть 1: баланс за каждый месяц ----------
    balance := initial
    for month := 1; month <= 12; month++ {
        balance = balance * (1 + monthly)
        fmt.Printf("месяц %2d: %.2f\n", month, balance)
    }

    // ---------- Часть 2: через сколько месяцев удвоение ----------
    balance = initial  // сброс, иначе считаем не с нуля
    months := 0
    for balance < initial*2 {  // for только с условием
        balance = balance * (1 + monthly)
        months++
    }

    fmt.Printf("Вклад удвоится через %d мес. Баланс: %.2f\n", months, balance)
}