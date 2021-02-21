package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	utils "practice/задачи_по_программированию"
	"time"
)

/*
Задача 1 Е. Степень
1 <= a <= 9
1 <= n <= 7000
*/
func main() {
	if len(os.Args) != 3 {
		usage()
		os.Exit(1)
	}

	a, n, err := input()
	if err != nil {
		log.Fatal(err)
	}

	if err := validate(a, n); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("a = %v, n = %v\n", a, n)

	runPow(a, n, powSTD)
	runPow(a, n, powSimple)
	runPow(a, n, pow2)
	runPow(a, n, powSmart)
}

func runPow(a, n uint64, pow func(uint64, uint64) uint64) {
	start := time.Now()
	res := pow(a, n)
	duration := time.Since(start)
	fmt.Printf("time %v, res = %v\n", duration, res)
}

func powSTD(a, n uint64) uint64 {
	return uint64(math.Pow(float64(a), float64(n)))
}

func powSimple(a, n uint64) (res uint64) {
	res = a
	for n--; n != 0; n-- {
		res *= a
	}
	return
}

func powSmart(a, n uint64) (res uint64) {
	if n == 1 || a == 1 {
		return a
	}

	more := n % 2

	a2 := a * a
	if n == 2 {
		return a2
	}

	res = powSmart(a2, n/2)

	if more == 1 {
		res *= a
	}
	return
}

func pow2(a, n uint64) uint64 {
	if n == 1 || a == 1 {
		return a
	}

	a2 := a * a
	if n == 2 {
		return a2
	}

	more := n % 2
	n = n / 2
	res := a2
	for n--; n != 0; n-- {
		res *= a2
	}

	if more == 1 {
		res *= a
	}

	return res
}

func usage() {
	fmt.Printf("%s <число> <степень>\n", os.Args[0])
}

func input() (a uint64, n uint64, err error) {
	a, err = utils.Str2uint64(os.Args[1])
	if err != nil {
		return
	}

	n, err = utils.Str2uint64(os.Args[2])
	return
}

func validate(a uint64, n uint64) error {
	if a < 1 || a > 9 {
		return errors.New("число должно быть в дипазоне [1; 9]")
	}

	if n < 1 || n > 7000 {
		return errors.New("степень должна быть в дипазоне [1; 7000]")
	}

	return nil
}
