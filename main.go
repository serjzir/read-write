package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	filename := "./output.txt"
	fo, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer fo.Close()

	filename_input := "./input.txt"
	f, err := os.OpenFile(filename_input, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileReader := bufio.NewReader(f)
	re := regexp.MustCompile(`(\d+)|\+|\-|\=`)
	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}

		res := re.FindAllString(string(line), -1)
		if len(res) > 0 {
			res0, _ := strconv.Atoi(res[0])
			res2, _ := strconv.Atoi(res[1])
			switch res[1] {
			case "+":
				result := res0 + res2
				_, _ = fo.WriteString(fmt.Sprint(res[0], res[1], res[2], res[3], result))
				_, _ = fo.WriteString("\n")
			case "-":
				result := res0 - res2
				_, _ = fo.WriteString(fmt.Sprint(res[0], res[1], res[2], res[3], result))
				_, _ = fo.WriteString("\n")
			}
		}
	}
}
