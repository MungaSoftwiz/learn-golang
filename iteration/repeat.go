package iteration

//const repeatCount = 5

func Repeat(character string, count int) string {
	//we use explicit version as we're not assign+init
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
