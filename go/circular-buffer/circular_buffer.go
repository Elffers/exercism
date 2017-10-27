package circular

import "fmt"

const testVersion = 4

type Buffer struct {
	size   int
	buffer string
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size:   size,
		buffer: "",
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if len(b.buffer) < 1 {
		return 0, fmt.Errorf("empty buffer")
	}

	out := b.buffer[0]
	b.buffer = b.buffer[1:]

	return out, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if len(b.buffer) >= b.size {
		return fmt.Errorf("buffer is full")
	}
	b.buffer += string(c)

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if len(b.buffer) >= b.size {
		b.buffer = b.buffer[1:]
	}

	b.buffer += string(c)
}

func (b *Buffer) Reset() {
	b.buffer = ""
}
