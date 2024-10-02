package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		usd = 15000
		eur = 16000
		jpy = 140
	)

	fmt.Print("Masukan nilai Rupiah: ")

	rupiahReader := bufio.NewReader(os.Stdin)
	rupiahStr, _ := rupiahReader.ReadString('\n')

	rupiahStr = strings.TrimSpace(rupiahStr)

	rupiah, err := strconv.ParseFloat(rupiahStr, 64)
	if err != nil {
		fmt.Println("Input tidak valid")
		return
	}

	fmt.Printf("Nilai USD: %.2f\n", rupiah/usd)
	fmt.Printf("Nilai EUR: %.2f\n", rupiah/eur)
	fmt.Printf("Nilai JPY: %.2f\n", rupiah/jpy)
}
