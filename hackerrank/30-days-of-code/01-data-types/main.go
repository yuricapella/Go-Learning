package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	j, _ := strconv.ParseUint(scanner.Text(), 10, 64)

	scanner.Scan()
	e, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	t := scanner.Text()

	fmt.Println(i + j)
	fmt.Printf("%.1f\n", d+e)
	fmt.Println(s + t)

}
