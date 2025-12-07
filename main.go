package main

import (
	"bufio"
	"fmt"
	"github.com/LavrikPavlov/go-learn/imt"
	"github.com/LavrikPavlov/go-learn/length_string"
	"github.com/LavrikPavlov/go-learn/types"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	getInfoCommand()
	var countError = 0
	var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	var reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nВыберете программу: ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var command, parseError = strconv.Atoi(input)
		if err == nil && parseError == nil {
			result := ""
			switch command {
			case 1:
				result = imt.ImtService()
			case 2:
				types.TypesService()
			case 3:
				length_string.StringService()
			case 0:
				exitApp()
			}
			if result != "" {
				logger.Printf("Программа вернула: %s", result)
			}
		} else {
			if countError == 3 {
				exitApp()
			}

			if parseError == nil {
				logger.Println(err)
			}

			if err == nil {
				logger.Println(parseError)
			}

			getInfoCommand()
			countError++
		}
	}
}

func getInfoCommand() {
	fmt.Println(`
	Команда [1] - сервси для расчета IMT
	Команда [2] - сервси для вывода инфы о всех переменных
	Команда [3] - сервси для игры в длину строки
	Команда [0] - выход
	`)
}

func exitApp() {
	os.Exit(0)
}
