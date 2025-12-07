package calc

/*
Разработка калькулятора
*/

func CalcService(value string) (result float64, err error) {
	var charArray = []rune(value)
	result = float64(charArray[0])
	return
}
