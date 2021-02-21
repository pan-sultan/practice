package utils

import "fmt"

func Str2uint64(s string) (n uint64, err error) {
	for _, ch := range s {
		if n != 0 {
			n *= 10
		}

		if ch < '0' || ch > '9' {
			err = fmt.Errorf("unkknow character %c", ch)
			return
		}

		n += uint64(ch - '0')
	}

	return
}
