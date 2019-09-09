package main
import (
	"fmt";
	"bufio";
	"os";
	"strings"
)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		str := scanner.Text();

		fmt.Println(strings.ToLower(str));
	}
}
