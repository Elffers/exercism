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
	m := FreqMap{}
	c := make(chan FreqMap, len(strings))

	go func() {
		for _, s := range strings {
			c <- Frequency(s)
		}
		close(c)
	}()

	for fmap := range c {
		for k, v := range fmap {
			m[k] += v
		}
	}
	return m
}
