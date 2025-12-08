package main

import (
	"bufio"
	"fmt"
	"github.com/LavrikPavlov/go-learn/imt"
	"github.com/LavrikPavlov/go-learn/length_string"
	"github.com/LavrikPavlov/go-learn/types"
	"log"
	"os"
	"os/exec"
	"runtime"
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
			case 4:
				types.TestTypesService()
			case 0:
				exitApp()
			default:
				getInfoCommand()
			}
			if result != "" {
				logger.Printf("\nПрограмма вернула: %s", result)
			}
			countError = 0
		} else {
			countError++
			if countError == 3 {
				exitApp()
				break
			}
			if parseError != nil {
				logger.Println(parseError)
				continue
			}
			if err != nil {
				logger.Println(err)
				continue
			}
			clearTerminal()
		}
	}
}

func getInfoCommand() {
	fmt.Println(`
	Команда [1] - сервси для расчета IMT
	Команда [2] - сервси для вывода инфы о всех переменных
	Команда [3] - сервси для игры в длину строки
	Команда [4] - сервис для теста типов
	Команда [0] - выход
	`)
}

func exitApp() {
	os.Exit(0)
}

func clearTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
