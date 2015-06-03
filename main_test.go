package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestClock(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanLines)
	main()
	fmt.Println("27")
	fmt.Println("0")
	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)
		result = append(result, x)
	}
	writer.WriteString("27\n")
	writer.WriteString("0\n")
	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println(x)
		result = append(result, x)
	}
}