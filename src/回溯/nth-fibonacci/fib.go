package nth_fibonacci

func fib(n int) int {
	return help(n)
}

func help(n int) int {
	if n == 1 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return help(n-1) + help(n-2)
	}
}
