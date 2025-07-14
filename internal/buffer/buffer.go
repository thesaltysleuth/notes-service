package buffer

// Buffer is a minimal byte accumulator
// It grows via append, mimicking bytes.Buffer internals

type Buffer struct {
	data []byte
}

func (b *Buffer) Write(p []byte) {
	b.data = append(b.data, p...)
}

func (b *Buffer) Bytes() []byte { return b.data }

