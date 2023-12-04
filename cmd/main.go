package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tip-peg-parser-go/pkg"
	"tip-peg-parser-go/pkg/tree"
)

func main() {
	code := "x = 10;\n" +
		"if (x > input)   { \n" +
		"	output x;\n" +
		"} else {\n" +
		"	output 0;\n" +
		"} \n"

	rootStatement := pkg.Parse(strings.Split(code, "\n"))
	builder := tree.ASTreeBuilder{Root: rootStatement}
	fmt.Printf("%s", builder.Build())
}

func openFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, err
}

func getFileLines(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	textBox := scanner.Text()
	lines := strings.Split(textBox, ",")
	return lines
}
