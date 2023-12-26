package culculation

import (
	"strconv"
)

func GetMaxNumber(a int) string {
	var c int
	b := strconv.Itoa(a)
	c = len(b)

	var y = string(b[0])
	for i := 0; i < c; i++ {
		if y <= string(b[i]) {
			y = string(b[i])
		}
	}
	return y
}
