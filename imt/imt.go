package imt

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func ImtService() string {
	var height float64
	var weight int

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите рост через точку в метрах и вес: ")
	weight = GetNum("Вес", INTEGER, reader).(int)
	height = GetNum("Рост в м.: ", FLOAT, reader).(float64)

	return calcUser(height, weight)
}

func calcUser(userHeight float64, userWeight int) string {
	var IMT = float64(userWeight) / math.Pow(userHeight, 2)
	fmt.Println("IMT пользователя: ", IMT)
	fmt.Printf("Результат %s", checkIMT(IMT))

	var kg, count, result = userWeight, userHeight, IMT
	return fmt.Sprintf("\nВес: %d | Рост:%.2f м. | Процент: %.2f%% \n", kg, count, result)

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
