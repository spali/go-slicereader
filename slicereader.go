package slicereader

import (
	"errors"
)

// EOS is the error returned by Read when no more element is available.
// Functions should return EOS only to signal a graceful end of input.
var EOS = errors.New("EOS")

type predicate = func(v interface{}) bool

// SliceReader supports reading a slice similar to an io.Reader.
type SliceReader struct {
	s []interface{}
	i int64
}

// NewSliceReader returns a new SliceReader.
func NewSliceReader(slice []interface{}) *SliceReader {
	return &SliceReader{
		s: slice,
		i: 0,
	}
}

// Len returns the number of the unread elements of the slice.
func (sr *SliceReader) Len() int {
	return int(sr.Size() - int64(sr.i))
}

// Size returns the original length of the underlying slice.
// The returned value is always the same and is not affected by calls
// to any other method.
func (sr *SliceReader) Size() int64 {
	return int64(len(sr.s))
}

// Read reads a single element of the slice or the EOS error will be returned.
func (sr *SliceReader) Read() (e interface{}, err error) {
	if sr.i >= sr.Size() {
		return nil, EOS
	}
	e = sr.s[sr.i]
	sr.i++
	return
}

// ReadWhile reads the slice till the element before the given predicate function returns false.
// If the end of the slice is reached the EOS error will be returned.
func (sr *SliceReader) ReadWhile(p predicate) (s []interface{}, err error) {
	s = make([]interface{}, 0)
	for sr.i < sr.Size() {
		if !p(sr.s[sr.i]) {
			return s, nil
		}
		s = append(s, sr.s[sr.i])
		sr.i++
	}
	return s, EOS
}

// ReadUntil reads the slice till the element before the given predicate function returns true.
// If the end of the slice is reached the EOS error will be returned.
func (sr *SliceReader) ReadUntil(p predicate) (s []interface{}, err error) {
	s = make([]interface{}, 0)
	for sr.i < sr.Size() {
		if p(sr.s[sr.i]) {
			return s, nil
		}
		s = append(s, sr.s[sr.i])
		sr.i++
	}
	return s, EOS
}

// ReadWhile reads the slice till including the element, the given predicate function returns false.
// If the end of the slice is reached the EOS error will be returned.
func (sr *SliceReader) ReadWhileIncl(p predicate) (s []interface{}, err error) {
	s = make([]interface{}, 0)
	for sr.i < sr.Size() {
		if !p(sr.s[sr.i]) {
			s = append(s, sr.s[sr.i])
			return s, nil
		}
		s = append(s, sr.s[sr.i])
		sr.i++
	}
	return s, EOS
}

// ReadUntil reads the slice till including  the element, before the given predicate function returns true.
// If the end of the slice is reached the EOS error will be returned.
func (sr *SliceReader) ReadUntilIncl(p predicate) (s []interface{}, err error) {
	s = make([]interface{}, 0)
	for sr.i < sr.Size() {
		if p(sr.s[sr.i]) {
			s = append(s, sr.s[sr.i])
			return s, nil
		}
		s = append(s, sr.s[sr.i])
		sr.i++
	}
	return s, EOS
}
