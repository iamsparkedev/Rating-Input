package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("here is a rating input program")
	fmt.Println("Enter your rating : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("here is your rating", num+10)
	}
}
