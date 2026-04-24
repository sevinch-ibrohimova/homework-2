package main

import (
	"alifmobi/task2/internal/receipt"
	"fmt"
	"strings"
)

func main() {
	var senderName, senderSurname string
	var receiverName, receiverSurname string
	var cardNumber string
	var summa, cardType int

	fmt.Print("Имя отправителя: ")
	fmt.Scan(&senderName)
	fmt.Print("Фамилия отправителя: ")
	fmt.Scan(&senderSurname)

	fmt.Print("Имя получателя: ")
	fmt.Scan(&receiverName)
	fmt.Print("Фамилия получателя: ")
	fmt.Scan(&receiverSurname)

	fmt.Print("Номер карты получателя (16 цифр): ")
	fmt.Scan(&cardNumber)

	fmt.Print("Введите сумму: ")
	fmt.Scan(&summa)

	if summa < 500 || summa > 15000000 {
		fmt.Println("Ошибка: лимит от 500 до 15 000 000 сум")
		return
	}

	fmt.Print("Alif карта? (1-да/0-нет): ")
	fmt.Scan(&cardType)

	senderFull := strings.ToUpper(senderSurname + " " + senderName)
	receiverFull := strings.ToUpper(receiverSurname + " " + receiverName)

	maskedCard := receipt.MaskCard(cardNumber)
	txID := receipt.GenerateTxID()
	dateTime := receipt.GetCurrentTime()
	komissiya := receipt.CalculateCommission(summa, cardType)
	itogo := summa + komissiya

	fmt.Println("\n====== Электронный чек ======")
	fmt.Printf("Отправитель:\t\t%s\n", senderFull)
	fmt.Printf("Получатель:\t\t%s\n", receiverFull)
	fmt.Printf("Номер транзакции:\t%d\n", txID)
	fmt.Printf("Счёт зачисления:\t%s\n", maskedCard)
	fmt.Printf("Дата и время:\t\t%s\n", dateTime)
	fmt.Printf("Сумма:\t\t\t%d сум\n", summa)
	fmt.Printf("Комиссия:\t\t%d сум\n", komissiya)
	fmt.Printf("Итого:\t\t\t%d сум\n", itogo)
	fmt.Println("АО ALIF TECH • Лицензия ЦБ РУз № 000010")
	fmt.Println("Статус:\t\t\tИсполнено")
	fmt.Println("============================")
}
