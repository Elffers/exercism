package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	res, err := o()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(o, input)
		}
		return

	}

	defer func() {
		if p := recover(); p != nil {
			if frob, ok := p.(FrobError); ok {
				res.Defrob(frob.defrobTag)
				err = frob.inner
			}
			err = p.(error)
		}
		res.Close()
	}()

	res.Frob(input)

	return err
}