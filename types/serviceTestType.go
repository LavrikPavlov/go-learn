package types

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestTypesService() {
	var reader = bufio.NewReader(os.Stdin)
	var logger = log.New(os.Stdout, "[types] ", log.LstdFlags)
	fmt.Println("=== Тестирование типов данных ===")
	isShowTypes := getAccept(reader, "Выбрать тип случайным образом? (y/n): ")

	var t reflect.StructField
	var err error

	if isShowTypes {
		t = randType()
		fmt.Printf("\nВыбран случайный тип: %s (%s)\n",
			t.Name,
			t.Tag.Get("описание"))
	} else {
		fmt.Println("\nДоступные типы:")
		showAvailableTypes()

		t, err = inType()
		if err != nil {
			logger.Println("Ошибка:", err)
			return
		}
	}

	fmt.Printf("\n=== Ввод значения для типа %s ===\n", t.Name)
	fmt.Println("Описание:", t.Tag.Get("описание"))

	value, err := inputValueWithRetry(reader, t.Name)
	if err != nil {
		logger.Printf("Ошибка ввода для типа %s: %v\n", t.Name, err)
		return
	}

	fmt.Printf("\n=== Результат ===\n")
	fmt.Printf("Тип: %s\n", t.Name)
	fmt.Printf("Введенное значение: %v\n", value)
	fmt.Printf("Go тип значения: %T\n", value)

	performTypeChecks(value, t.Name)
}

func inputValueWithRetry(reader *bufio.Reader, needType string) (interface{}, error) {
	maxAttempts := 3

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		fmt.Printf("\nПопытка %d/%d\n", attempt, maxAttempts)
		fmt.Printf("Введите значение для типа %s: ", needType)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := convertStringToType(input, needType)
		if err == nil {
			return value, nil
		}

		fmt.Printf("Ошибка: %v\n", err)
		showTypeHint(needType)

		if attempt < maxAttempts {
			fmt.Println("Попробуйте еще раз.")
		}
	}

	return nil, fmt.Errorf("превышено максимальное количество попыток (%d)", maxAttempts)
}

func convertStringToType(input, needType string) (interface{}, error) {
	switch needType {
	case "Int8":
		val, err := strconv.ParseInt(input, 10, 8)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для int8: %v. Диапазон: -128 до 127", err)
		}
		return int8(val), nil

	case "Int16":
		val, err := strconv.ParseInt(input, 10, 16)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для int16: %v. Диапазон: -32768 до 32767", err)
		}
		return int16(val), nil

	case "Int32":
		val, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для int32: %v", err)
		}
		return int32(val), nil

	case "Int64":
		val, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для int64: %v", err)
		}
		return int64(val), nil

	case "Int":
		val, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для int: %v", err)
		}
		return int(val), nil

	case "Uint8":
		val, err := strconv.ParseUint(input, 10, 8)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для uint8: %v. Диапазон: 0 до 255", err)
		}
		return uint8(val), nil

	case "Uint16":
		val, err := strconv.ParseUint(input, 10, 16)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для uint16: %v", err)
		}
		return uint16(val), nil

	case "Uint32":
		val, err := strconv.ParseUint(input, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для uint32: %v", err)
		}
		return uint32(val), nil

	case "Uint64":
		val, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для uint64: %v", err)
		}
		return uint64(val), nil

	case "Uint":
		val, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для uint: %v", err)
		}
		return uint(val), nil

	case "Float32":
		val, err := strconv.ParseFloat(input, 32)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для float32: %v", err)
		}
		return float32(val), nil

	case "Float64":
		val, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное значение для float64: %v", err)
		}
		return val, nil

	case "String":
		if input == "" {
			return nil, fmt.Errorf("строка не может быть пустой")
		}
		return input, nil

	case "Bool":
		switch strings.ToLower(input) {
		case "true", "1", "yes", "y", "да", "истина":
			return true, nil
		case "false", "0", "no", "n", "нет", "ложь":
			return false, nil
		default:
			return nil, fmt.Errorf("неверное значение для bool. Используйте: true/false, yes/no, 1/0, да/нет")
		}

	case "Rune":
		if len(input) != 1 {
			return nil, fmt.Errorf("rune должен быть одним символом")
		}
		return []rune(input)[0], nil

	case "Byte":
		if len(input) != 1 {
			return nil, fmt.Errorf("byte должен быть одним символом")
		}
		return input[0], nil

	default:
		return nil, fmt.Errorf("тип %s не поддерживается для ввода", needType)
	}
}

func showTypeHint(needType string) {
	switch needType {
	case "Int8":
		fmt.Println("Подсказка: int8 принимает числа от -128 до 127")
	case "Int16":
		fmt.Println("Подсказка: int16 принимает числа от -32768 до 32767")
	case "Uint8":
		fmt.Println("Подсказка: uint8 принимает числа от 0 до 255")
	case "Bool":
		fmt.Println("Подсказка: используйте true/false, yes/no, 1/0, да/нет")
	case "Rune", "Byte":
		fmt.Println("Подсказка: введите один символ")
	case "String":
		fmt.Println("Подсказка: строка не может быть пустой")
	}
}

func performTypeChecks(value interface{}, typeName string) {
	fmt.Println("\n=== Проверки ===")

	switch v := value.(type) {
	case int8:
		if v == 0 {
			fmt.Println("Предупреждение: значение равно 0")
		}
		if v == 127 {
			fmt.Println("Внимание: достигнут максимальный предел int8")
		}
		if v == -128 {
			fmt.Println("Внимание: достигнут минимальный предел int8")
		}

	case int16:
		if v > 10000 {
			fmt.Println("Значение больше 10000")
		}

	case int32:
		if typeName == "Rune" {
			fmt.Printf("Символ: %c\n", v)
			if v >= 'A' && v <= 'Z' {
				fmt.Println("Это заглавная латинская буква")
			} else if v >= 'a' && v <= 'z' {
				fmt.Println("Это строчная латинская буква")
			} else if v >= '0' && v <= '9' {
				fmt.Println("Это цифра")
			}
		}

	case uint8:
		if typeName == "Byte" {
			fmt.Printf("Символ ASCII: %c\n", v)
			if v >= 32 && v <= 126 {
				fmt.Println("Это печатный символ ASCII")
			}
		}
	case string:
		fmt.Printf("Длина строки: %d символов\n", len(v))
		if strings.Contains(v, " ") {
			fmt.Println("Строка содержит пробелы")
		}
	case bool:
		if v {
			fmt.Println("Значение истинно")
		} else {
			fmt.Println("Значение ложно")
		}
	case float32, float64:
		fmt.Printf("Научная запись: %e\n", v)
	}
}

func randType() reflect.StructField {
	allTypes := AllTypes{}
	t := reflect.TypeOf(allTypes)

	numFields := t.NumField()
	if numFields == 0 {
		return reflect.StructField{}
	}

	index := rand.Intn(numFields)
	return t.Field(index)
}

func inType() (reflect.StructField, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nВведите имя типа (или 'list' для списка, 'exit' для выхода): ")

		inStr, _ := reader.ReadString('\n')
		inStr = strings.TrimSpace(inStr)

		if strings.EqualFold(inStr, "list") {
			showAvailableTypes()
			continue
		}

		if strings.EqualFold(inStr, "exit") {
			return reflect.StructField{}, fmt.Errorf("выход по запросу пользователя")
		}

		allTypes := AllTypes{}
		t := reflect.TypeOf(allTypes)

		field, found := t.FieldByName(inStr)
		if !found {
			for i := 0; i < t.NumField(); i++ {
				f := t.Field(i)
				if strings.EqualFold(f.Name, inStr) {
					field = f
					found = true
					break
				}
			}
		}

		if !found {
			fmt.Printf("Тип '%s' не найден. ", inStr)
			fmt.Println("Используйте 'list' для просмотра доступных типов.")
			continue
		}

		return field, nil
	}
}
