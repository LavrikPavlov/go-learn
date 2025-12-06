package imt

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/LavrikPavlov/go-learn/imt/functions"
)

func ImtService() {
	var height float64
	var weight int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите рост через точку в метрах и вес: ")
	weight = functions.GetNum("Вес", functions.INTEGER, reader).(int)
	height = functions.GetNum("Рост в м.: ", functions.FLOAT, reader).(float64)

	calcUser(height, weight)
}

func calcUser(userHeight float64, userWeight int) {
	var IMT = float64(userWeight) / math.Pow(userHeight, 2)
	fmt.Println("IMT пользователя: ", IMT)
	fmt.Printf("Результат %s", checkIMT(IMT))

	var kg, count = userWeight, userHeight
	fmt.Printf("\nВес: %d | Рост:%.1f м.\n", kg, count)

}

func checkIMT(imt float64) string {
	switch {
	case imt < 16:
		return "Выраженный дефицит массы тела"
	case imt >= 16 && imt < 18.5:
		return "Недостаточная (дефицит) масса тела"
	case imt >= 18.5 && imt < 25:
		return "Норма"
	case imt >= 25 && imt < 30:
		return "Избыточная масса тела (предожирение)"
	case imt >= 30 && imt < 35:
		return "Ожирение первой степени"
	case imt >= 35 && imt < 40:
		return "Ожирение второй степени"
	default:
		return "Ожирение третьей степени (морбидное)"
	}
}
