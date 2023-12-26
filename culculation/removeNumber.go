package culculation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func RemoveTheNumber(a, b int) string {
	var d int
	c := strconv.Itoa(a)
	d = len(c)

	var k []string
	var y string

	for i := 0; i < d; i++ {
		y = string(c[i])
		if y != strconv.Itoa(b) {
			k = append(k, y)
		}

	}
	var f = strings.Join(k, "")
	//fmt.Println(f)
	if c == f {
		err := errors.New("Искомое число отсутствует")
		fmt.Println("", err)
	} else {
		return f
	}

	return f
}
