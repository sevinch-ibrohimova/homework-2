package receipt

import (
	"math/rand"
	"time"
)

func MaskCard(cardNumber string) string {
	last4 := cardNumber[len(cardNumber)-4:]
	return "UZCARD**" + last4
}

func GenerateTxID() int {
	return rand.Intn(900000000) + 100000000
}

func GetCurrentTime() string {
	// 02.01.2006 15:04 - это магическая дата-шаблон в Go
	return time.Now().Format("02.01.2006 15:04")
}

func CalculateCommission(summa int, cardType int) int {
	if cardType == 1 {
		return 0
	}
	return int(float64(summa) * 0.0029)
}
