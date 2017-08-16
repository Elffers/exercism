package letter

const testVersion = 1

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	c := make(chan FreqMap, len(strings))

	for _, s := range strings {
		go func(in string) {
			c <- Frequency(in)
		}(s)
	}

	m := FreqMap{}

	for i := 0; i < len(strings); i++ {
		fmap := <-c
		for k, v := range fmap {
			m[k] += v
		}
	}
	return m
}
