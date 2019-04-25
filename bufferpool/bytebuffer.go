package bufferpool

import "io"

// ByteBuffer provides byte buffer, which can be used for minimizing
// memory allocations.
//
// ByteBuffer may be used with functions appending data to the given []byte
// slice. See example code for details.
//
// Use Get for obtaining an empty byte buffer.
type ByteBuffer struct {
  // b is a byte buffer to use in append-like workloads.
  // See example code for details.
  b []byte
}

// Len returns the size of the byte buffer.
func (b *ByteBuffer) Len() int {
  return len(b.b)
}

// ReadFrom implements io.ReaderFrom.
//
// The function appends all the data read from r to b.
func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) {
  p := b.b
  nStart := int64(len(p))
  nMax := int64(cap(p))
  n := nStart
  if nMax == 0 {
    nMax = 64
    p = make([]byte, nMax)
  } else {
    p = p[:nMax]
  }
  for {
    if n == nMax {
      nMax *= 2
      bNew := make([]byte, nMax)
      copy(bNew, p)
      p = bNew
    }
    nn, err := r.Read(p[n:])
    n += int64(nn)
    if err != nil {
      b.b = p[:n]
      n -= nStart
      if err == io.EOF {
        return n, nil
      }
      return n, err
    }
  }
}

// WriteTo implements io.WriterTo.
func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) {
  n, err := w.Write(b.b)
  return int64(n), err
}

// Bytes returns b.b, i.e. all the bytes accumulated in the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
func (b *ByteBuffer) Bytes() []byte {
  return b.b
}

// Write implements io.Writer - it appends p to ByteBuffer.b
func (b *ByteBuffer) Write(p []byte) (int, error) {
  b.b = append(b.b, p...)
  return len(p), nil
}

// WriteByte appends the byte c to the buffer.
//
// The purpose of this function is bytes.Buffer compatibility.
//
// The function always returns nil.
func (b *ByteBuffer) WriteByte(c byte) error {
  b.b = append(b.b, c)
  return nil
}

// WriteString appends s to ByteBuffer.b.
func (b *ByteBuffer) WriteString(s string) (int, error) {
  b.b = append(b.b, s...)
  return len(s), nil
}

// Set sets ByteBuffer.b to p.
func (b *ByteBuffer) Set(p []byte) {
  b.b = append(b.b[:0], p...)
}

// SetString sets ByteBuffer.b to s.
func (b *ByteBuffer) SetString(s string) {
  b.b = append(b.b[:0], s...)
}

// String returns string representation of ByteBuffer.b.
func (b *ByteBuffer) String() string {
  return string(b.b)
}

// Reset makes ByteBuffer.b empty.
func (b *ByteBuffer) Reset() {
  b.b = b.b[:0]
}

// Reset makes ByteBuffer.b empty.
func (b *ByteBuffer) Release() {
  b.Reset()
  Put(b)
}
