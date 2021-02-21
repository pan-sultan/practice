package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	utils "practice/задачи_по_программированию"
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

	fmt.Printf("%v %v\n", a, n)
	fmt.Println(pow(a, n))
}

func pow(a, n uint64) (res uint64) {
	res = a
	for n--; n != 0; n-- {
		res *= a
	}
	return
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
