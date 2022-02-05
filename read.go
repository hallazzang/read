package read

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	_ Reader = (*UnbufferedReader)(nil)
	_ Reader = (*BufferedReader)(nil)
)

func New(r io.Reader) *BufferedReader {
	return NewBuffered(r)
}

type Reader interface {
	Bool() (bool, error)
	Int() (int, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)
	Uint() (uint, error)
	Uint8() (uint8, error)
	Uint16() (uint16, error)
	Uint32() (uint32, error)
	Uint64() (uint64, error)
	Float32() (float32, error)
	Float64() (float64, error)
	Complex64() (complex64, error)
	Complex128() (complex128, error)
	String() (string, error)
	Line() (string, error)
}

func newStdinReader() *UnbufferedReader {
	return &UnbufferedReader{os.Stdin}
}

func Bool() (bool, error) {
	return newStdinReader().Bool()
}

func Int() (int, error) {
	return newStdinReader().Int()
}

func Int8() (int8, error) {
	return newStdinReader().Int8()
}

func Int16() (int16, error) {
	return newStdinReader().Int16()
}

func Int32() (int32, error) {
	return newStdinReader().Int32()
}

func Int64() (int64, error) {
	return newStdinReader().Int64()
}

func Uint() (uint, error) {
	return newStdinReader().Uint()
}

func Uint8() (uint8, error) {
	return newStdinReader().Uint8()
}

func Uint16() (uint16, error) {
	return newStdinReader().Uint16()
}

func Uint32() (uint32, error) {
	return newStdinReader().Uint32()
}

func Uint64() (uint64, error) {
	return newStdinReader().Uint64()
}

func Float32() (float32, error) {
	return newStdinReader().Float32()
}

func Float64() (float64, error) {
	return newStdinReader().Float64()
}

func Complex64() (complex64, error) {
	return newStdinReader().Complex64()
}

func Complex128() (complex128, error) {
	return newStdinReader().Complex128()
}

func String() (string, error) {
	return newStdinReader().String()
}

func Line() (string, error) {
	return newStdinReader().Line()
}

type UnbufferedReader struct {
	io.Reader
}

func NewUnbuffered(r io.Reader) *UnbufferedReader {
	return &UnbufferedReader{r}
}

func (r *UnbufferedReader) Bool() (b bool, err error) {
	_, err = fmt.Fscan(r, &b)
	return
}

func (r *UnbufferedReader) Int() (i int, err error) {
	_, err = fmt.Fscan(r, &i)
	return
}

func (r *UnbufferedReader) Int8() (i int8, err error) {
	_, err = fmt.Fscan(r, &i)
	return
}

func (r *UnbufferedReader) Int16() (i int16, err error) {
	_, err = fmt.Fscan(r, &i)
	return
}

func (r *UnbufferedReader) Int32() (i int32, err error) {
	_, err = fmt.Fscan(r, &i)
	return
}

func (r *UnbufferedReader) Int64() (i int64, err error) {
	_, err = fmt.Fscan(r, &i)
	return
}

func (r *UnbufferedReader) Uint() (u uint, err error) {
	_, err = fmt.Fscan(r, &u)
	return
}

func (r *UnbufferedReader) Uint8() (u uint8, err error) {
	_, err = fmt.Fscan(r, &u)
	return
}

func (r *UnbufferedReader) Uint16() (u uint16, err error) {
	_, err = fmt.Fscan(r, &u)
	return
}

func (r *UnbufferedReader) Uint32() (u uint32, err error) {
	_, err = fmt.Fscan(r, &u)
	return
}

func (r *UnbufferedReader) Uint64() (u uint64, err error) {
	_, err = fmt.Fscan(r, &u)
	return
}

func (r *UnbufferedReader) Float32() (f float32, err error) {
	_, err = fmt.Fscan(r, &f)
	return
}

func (r *UnbufferedReader) Float64() (f float64, err error) {
	_, err = fmt.Fscan(r, &f)
	return
}

func (r *UnbufferedReader) Complex64() (c complex64, err error) {
	_, err = fmt.Fscan(r, &c)
	return
}

func (r *UnbufferedReader) Complex128() (c complex128, err error) {
	_, err = fmt.Fscan(r, &c)
	return
}

func (r *UnbufferedReader) String() (s string, err error) {
	_, err = fmt.Fscan(r, &s)
	return
}

func (r *UnbufferedReader) Line() (string, error) {
	var buf strings.Builder
	b := make([]byte, 1)
	var prev byte
	for {
		if _, err := r.Read(b); err != nil {
			return "", err
		}
		_, _ = buf.Write(b)
		if b[0] == '\n' {
			break
		}
		prev = b[0]
	}
	if prev == '\r' {
		return strings.TrimSuffix(buf.String(), "\r\n"), nil
	}
	return strings.TrimSuffix(buf.String(), "\n"), nil
}

type BufferedReader struct {
	UnbufferedReader
}

func NewBuffered(r io.Reader) *BufferedReader {
	return &BufferedReader{UnbufferedReader{bufio.NewReader(r)}}
}

func (r *BufferedReader) Line() (s string, err error) {
	br := r.Reader.(*bufio.Reader)
	var chunks [][]byte
	for {
		b, isPrefix, err := br.ReadLine()
		if err != nil {
			return "", err
		}
		chunks = append(chunks, b)
		if !isPrefix {
			break
		}
	}
	return string(bytes.Join(chunks, []byte{})), nil
}
