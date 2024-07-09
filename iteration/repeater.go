package iteration

func Repeat(value string, iterations int) string {
	var variable string
	for i := 0; i < iterations; i++ {
		variable += value
	}
	return variable
}
