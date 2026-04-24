package main

import (
	"alifmobi/internal/commission"
	"fmt"
)

func main() {
	var summa int
	var cardType int
	fmt.Print("Введите сумму: ")
	fmt.Scan(&summa)
	
	if summa < 500 || summa > 15000000 {
		fmt.Println("Ошибка: лимит от 500 до 15 000 000 сум")
		return
	}

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&cardType)
	komissiya, itogo := commission.Calculate(summa, cardType)

	fmt.Println("\n======== Чек ========")
	fmt.Println("Услуга: SHOKIROV SHUHRAT")
	fmt.Printf("Сумма: %d сум\n", summa)
	fmt.Printf("Комиссия: %d сум\n", komissiya)
	fmt.Printf("Итого: %d сум\n", itogo)
	fmt.Println("Статус: Исполнено")
}
