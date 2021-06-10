package task7

import (
	"fmt"
)

func RuneByIndex(s *string, i *int) (runa rune, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}

	}()
	runa = rune((*s)[*i])
	return runa, err
}
