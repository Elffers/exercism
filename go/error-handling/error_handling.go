package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	res, _ := o()

	res.Close()
	res.Frob(input)

	return
}
