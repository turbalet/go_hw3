package task3

import (
	"errors"
)

func Atoi(s string) (int, error) {
	var res int
	for _, r := range []byte(s) {
		if r-'0' > 9 {
			return 0, errors.New("atoi: unable to cast")
		}
		res = res*10 + int(r-'0')
	}
	return res, nil
}
