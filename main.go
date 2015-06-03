package main

import (
	"bufio"
	"fmt"
	. "github.com/nickstew/golang-rolling-ball-clock/clock"
	"os"
	"strconv"
)

func ReadClocks() ([]int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		if x == 0 {
			break
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	fmt.Println("Please enter starting balls in queue for Ball Clock then press enter.\n",
		"Valid values are (27-127).\n",
		"Enter 0 to begin calculation of previously entered Ball Clocks.")
	clocks, err := ReadClocks()
	if err != nil {
		fmt.Println("You did not enter a valid integer between 27 and 127.  Please rerun program.")
	} else {
		for _, v := range clocks {
			clock := New(v)
			days := clock.FindCycleDays()
			fmt.Println(strconv.Itoa(v) + " balls cycle after " + strconv.Itoa(days) + " days.")
		}
	}
}
