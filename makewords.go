package main
import (
	"fmt";
	"bufio";
	"os";
	"strings"
)
func main() {
    reader := bufio.NewReader(os.Stdin);
    text, _ := reader.ReadString('\n');
    words := strings.Split(text, " ");
    
    for i := 0; i < len(words); i++ {
        word := words[i];
        wordWithoutDot := strings.Replace(word, ".", "", -1);
        wordWithoutComma := strings.Replace(wordWithoutDot, ",", "", -1);
    	fmt.Println(wordWithoutComma);
    }
}
