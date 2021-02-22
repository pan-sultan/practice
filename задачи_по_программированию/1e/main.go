package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	utils "practice/задачи_по_программированию"
	"time"
)

/*
Задача 1 Е. Степень
Вычислить степень (n) числа (a):
1 <= a <= 9
1 <= n <= 7000
*/

var one *big.Int

func init() {
	one = new(big.Int)
	one.SetUint64(1)
}

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

	runPow(a, n, powSimple)
	runPow(a, n, pow2)
	runPow(a, n, powSmartRec)
	runPow(a, n, powSmart)
}

func runPow(a *big.Int, n uint64, pow func(*big.Int, uint64) *big.Int) {
	start := time.Now()
	res := pow(a, n)
	duration := time.Since(start)
	fmt.Printf("time %v, res = %v\n", duration, res)
}

func powSimple(a *big.Int, n uint64) (res *big.Int) {
	res = new(big.Int)
	res.Set(a)
	for n--; n != 0; n-- {
		res = res.Mul(res, a)
	}
	return
}

func powSmartRec(a *big.Int, n uint64) *big.Int {
	var calc func(a big.Int, n uint64) big.Int

	calc = func(a big.Int, n uint64) (res big.Int) {
		if n == 1 || a.Cmp(one) == 0 {
			return a
		}

		more := n % 2
		var a2 big.Int
		a2.Set(&a)
		a2 = *a2.Mul(&a2, &a)
		if n == 2 {
			return a2
		}

		res = calc(a2, n/2)

		if more == 1 {
			res = *res.Mul(&res, &a)
		}
		return res
	}

	res := calc(*a, n)
	return &res
}

func powSmart(a *big.Int, n uint64) *big.Int {
	if n == 1 || a.Cmp(one) == 0 {
		return a
	}

	pows := make([]*big.Int, 0, 10)
	res := new(big.Int)
	res.Set(a)

	for ; n > 1; n /= 2 {
		if (n % 2) == 1 {
			tmp := new(big.Int)
			tmp.Set(res)
			pows = append(pows, tmp)
		}
		res = res.Mul(res, res)
	}

	for _, p := range pows {
		res = res.Mul(res, p)
	}

	return res
}

func pow2(a *big.Int, n uint64) *big.Int {
	if n == 1 || a.Cmp(one) == 0 {
		return a
	}

	a2 := new(big.Int)
	a2.Set(a)
	a2 = a2.Mul(a2, a)
	if n == 2 {
		return a2
	}

	more := n % 2
	n = n / 2
	res := new(big.Int)
	res.Set(a2)
	for n--; n != 0; n-- {
		res = res.Mul(res, a2)
	}

	if more == 1 {
		res = res.Mul(res, a)
	}

	return res
}

func usage() {
	fmt.Printf("%s <число> <степень>\n", os.Args[0])
}

func input() (*big.Int, uint64, error) {
	a, err := utils.Str2uint64(os.Args[1])
	if err != nil {
		return nil, 0, err
	}

	n, err := utils.Str2uint64(os.Args[2])
	if err != nil {
		return nil, 0, err
	}

	ba := new(big.Int)
	ba.SetUint64(a)
	return ba, n, nil
}

func validate(a *big.Int, n uint64) error {
	if a.Uint64() < 1 || a.Uint64() > 9 {
		return errors.New("число должно быть в дипазоне [1; 9]")
	}

	if n < 1 || n > 7000 {
		return errors.New("степень должна быть в дипазоне [1; 7000]")
	}

	return nil
}
