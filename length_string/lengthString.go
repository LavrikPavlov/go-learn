package length_string

import "fmt"

func StringService() {
	var message = ""
	var maxLength = genNum(15)
	fmt.Println("Максимальная длина сгенерирована")
	fmt.Print("Введите слово и нажмите Enter: ")
	_, _ = fmt.Scanln(&message)

	fmt.Printf("Ваше слово [%s]\n", message)

	length := len(message)
	fmt.Printf("Длина вашего слово %d \n", length)

	if maxLength >= length {
		fmt.Println("Вы попали!!!")
	} else {
		fmt.Println("Проиграл")
	}
}
