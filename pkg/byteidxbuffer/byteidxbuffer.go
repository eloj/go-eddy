package byteidxbuffer

// ByteIdxBuffer represents a buffer of variable-length members, and an indexing-array to access them.
// For example useful for cache-efficient storage of strings.
type ByteIdxBuffer struct {
	buffer []byte
	idx    []uint32
}

// NewByteIdxBuffer creates a new buffer and returns a pointer to it.
func NewByteIdxBuffer() *ByteIdxBuffer {
	res := &ByteIdxBuffer{
		buffer: make([]byte, 0),
		idx:    make([]uint32, 1),
	}

	return res
}

func (bib *ByteIdxBuffer) Len() int {
	return len(bib.idx) - 1
}

func (bib *ByteIdxBuffer) AddString(data string) {
	packet := []byte(data)
	bib.buffer = append(bib.buffer, packet...)
	bib.idx = append(bib.idx, uint32(len(bib.buffer)))
}

func (bib *ByteIdxBuffer) GetString(idx int) string {
	if idx >= len(bib.idx) || idx < 0 {
		panic("Out of bounds")
	}
	return string(bib.buffer[bib.idx[idx]:bib.idx[idx+1]])
}

func (bib *ByteIdxBuffer) GetSlice(idx int) []byte {
	if idx >= len(bib.idx) || idx < 0 {
		panic("Out of bounds")
	}
	return bib.buffer[bib.idx[idx]:bib.idx[idx+1]]
}
