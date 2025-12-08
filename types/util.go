package types

import (
	"bufio"
	"fmt"
	"strings"
)

func getAccept(reader *bufio.Reader, text string) bool {
	fmt.Printf("\n %s (yes/no) default no: ", text)
	answer, _ := reader.ReadString('\n')
	return strings.Contains(strings.ToLower(answer), "yes")
}
