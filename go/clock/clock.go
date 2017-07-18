package clock

import "fmt"

const testVersion = 4

type Clock int

func New(hour, minute int) Clock {
	minutes := mod((hour*60+minute), 1440)
	return Clock(minutes)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c /60, c % 60)
}

func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+ minutes)
}

func mod(n int, d int) int {
	r := n % d
	if r < 0 {
		r += d
	}
	return r
}
