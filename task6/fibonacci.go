package task6

func Fibonacci() func() int {
	n1, n2 := 0, 1
	return func() int {
		res := n1
		n1, n2 = n2, n2+n1
		return res
	}
}
