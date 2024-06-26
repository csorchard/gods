package a_bitmap

import (
	"bytes"
	"fmt"
)

type Bitmap struct {
	words  []uint64 // each bit of words[i] represents a group of 64 numbers
	length int
}

func New(size uint32) *Bitmap {
	newBitmap := &Bitmap{}
	word := int(size / 64)
	for word >= len(newBitmap.words) {
		newBitmap.words = append(newBitmap.words, 0)
	}
	return newBitmap
}

func (bitmap *Bitmap) Has(num uint32) bool {
	word, bit := num/64, uint(num%64)
	return int(word) < len(bitmap.words) && (bitmap.words[word]&(1<<bit)) != 0
}

func (bitmap *Bitmap) Add(num uint32) {
	word, bit := num/64, uint(num%64)
	for int(word) >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
		bitmap.length++
	}
}

func (bitmap *Bitmap) Len() int {
	return bitmap.length
}

func (bitmap *Bitmap) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, v := range bitmap.words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < 64; j++ {
			if v&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				_, _ = fmt.Fprintf(&buf, "%d", 64*uint(i)+j)
			}
		}
	}
	buf.WriteByte('}')
	_, _ = fmt.Fprintf(&buf, "\nLength: %d", bitmap.length)
	return buf.String()
}
