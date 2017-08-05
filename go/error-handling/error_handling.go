package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) error {
	res, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
	}

	res.Close()
	res.Frob(input)

	return nil
}
