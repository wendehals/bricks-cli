package utils

func Identity(k string) string {
	return k
}

func Add(x1, x2 int) int {
	return x1 + x2
}

func Subtract(x1, x2 int) int {
	return x1 - x2
}

func Max(x1, x2 int) int {
	if x1 >= x2 {
		return x1
	}
	return x2
}
