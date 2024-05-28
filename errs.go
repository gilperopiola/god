package god

func IfErrThen(err error, do func()) {
	if err != nil {
		do()
	}
}
