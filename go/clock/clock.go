package clock

import "fmt"

const testVersion = 4

type Clock struct{h, m int }

func New(hour, minute int) Clock {
	var m, overflow int
	overflow, m = divmod(minute, 60)

	h := mod(hour + overflow, 24)

	if minute < 0 {
		h = mod(h-1, 24)
	}

	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	raw_minutes := c.m + minutes

	var m, overflow int
	overflow, m = divmod(raw_minutes, 60)

	h := mod(c.h + overflow, 24)

	if raw_minutes < 0 && m != 0 { // this is mysterious
		h = mod(h-1, 24)
	}

	c.m = m
	c.h = h

	return c
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func mod(n int, d int) int {
	r := n % d
	if r < 0 {
		r += d
	}
	return r
}

func divmod(n int, d int) (int, int) {
	q := n / d
	r := mod(n, d)

	return q, r
}
