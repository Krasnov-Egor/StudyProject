package culculation

func GetReverseOrder(a int) []int {
	var k []int

	for a > 0 {
		var b int = a % 10
		k = append(k, b)
		a = a / 10
	}
	return k

}
