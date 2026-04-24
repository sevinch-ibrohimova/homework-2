package commission

func Calculate(summa int, cardType int) (int, int) {
	var komissiya int
	if cardType == 1 {
		komissiya = 0
	} else {
		komissiya = int(float64(summa) * 0.0029)
	}
	itogo := summa + komissiya
	return komissiya, itogo
}
