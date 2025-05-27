package utils

func Check(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
