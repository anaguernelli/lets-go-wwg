package main
import (
	“os”
	“bufio”
)

func leiaLinha() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
//incompleto