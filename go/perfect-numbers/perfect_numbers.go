package perfect

import "errors"

type Classification string

const (
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationPerfect   Classification = "ClassificationPerfect"
	ClassificationAbundant  Classification = "ClassificationAbundant"
)

const testVersion = 1

var ErrOnlyPositive = errors.New("Can't classify zero or negative.")

func Classify(n uint64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	sum := aliquotSum(int(n))
	if sum == int(n) {
		return ClassificationPerfect, nil
	}
	if sum > int(n) {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}

func aliquotSum(n int) (sum int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return
}
