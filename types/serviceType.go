package types

import (
	"bufio"
	"os"
)

func TypesService() {
	var reader = bufio.NewReader(os.Stdin)
	isShowTypes := getAccept(reader, "Показать все типы?")

	if isShowTypes {
		showAvailableTypes()
	}
}
