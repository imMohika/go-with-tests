package iteration

func Repeat(str string, count int) string {
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += str
	}
	return repeated
	//return strings.Repeat(s, 5)
}
