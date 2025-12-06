package functions

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type TypeValue string

const (
	FLOAT   TypeValue = "FLOAT"
	STRING  TypeValue = "STRING"
	INTEGER TypeValue = "INTEGER"
)

func GetNum(value string, typeValue TypeValue, reader *bufio.Reader) interface{} {
	var writer io.Writer = os.Stdout
	var logger = log.New(writer, "", log.LstdFlags)
	var count int = 0
	for {
		fmt.Printf("\n%s: ", value)
		inStr, _ := reader.ReadString('\n')
		inStr = strings.TrimSpace(inStr)

		result, err := getValue(typeValue, inStr)

		if err == nil {
			return result
		} else {
			logger.Printf("Ошибочное значение: %s [%v]\n", inStr, err)
			count++
		}

		if count == 10 {
			logger.Println("Кол-во попыток закончилось, принудительное завершение")
			os.Exit(0)
		}
	}
}

func getValue(typeValue TypeValue, input string) (interface{}, error) {
	switch typeValue {
	case FLOAT:
		return strconv.ParseFloat(input, 64)
	case INTEGER:
		return strconv.Atoi(input)
	case STRING:
		return input, nil
	default:
		return nil, fmt.Errorf("неизвестный тип: %s", typeValue)
	}
}
