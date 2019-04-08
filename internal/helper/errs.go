package helper

// RaiseErr if errs != nil
func RaiseErr(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
