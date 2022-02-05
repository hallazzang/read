package read_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hallazzang/read"
)

func mockStdin(content []byte) (*os.File, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("pipe: %w", err)
	}
	defer w.Close()
	if _, err := w.Write(content); err != nil {
		return nil, fmt.Errorf("write: %w", err)
	}
	return r, nil
}

func TestBasic(t *testing.T) {
	r, err := mockStdin([]byte(" 1  2 \n 3 4 5 \n 1 2\n"))
	require.NoError(t, err)
	defer r.Close()
	os.Stdin = r

	i, err := read.Int()
	require.NoError(t, err)
	require.Equal(t, 1, i)

	i, err = read.Int()
	require.NoError(t, err)
	require.Equal(t, 2, i)

	s, err := read.String()
	require.NoError(t, err)
	require.Equal(t, "3", s)

	s, err = read.Line()
	require.NoError(t, err)
	require.Equal(t, "4 5 ", s)

	i, err = read.Int()
	require.NoError(t, err)
	require.Equal(t, 1, i)

	s, err = read.Line()
	require.NoError(t, err)
	require.Equal(t, "2", s)
}

func TestUnbufferedReader_Line(t *testing.T) {
	r := read.NewUnbuffered(bytes.NewReader([]byte(" foo \n")))
	s, err := r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo ", s)

	r = read.NewUnbuffered(bytes.NewReader([]byte(" foo \r\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo ", s)

	r = read.NewUnbuffered(bytes.NewReader([]byte(" foo \t\r\r\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo \t\r", s)

	r = read.NewUnbuffered(bytes.NewReader([]byte(" foo \nbar\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo ", s)
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, "bar", s)

	r = read.NewUnbuffered(bytes.NewReader([]byte("\nfoo\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, "", s)
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, "foo", s)
}

func TestBufferedReader_Line(t *testing.T) {
	r := read.NewBuffered(bytes.NewReader([]byte(" foo \n")))
	s, err := r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo ", s)

	r = read.NewBuffered(bytes.NewReader([]byte(" foo \r\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo ", s)

	r = read.NewBuffered(bytes.NewReader([]byte(" foo \t\r\r\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo \t\r", s)

	r = read.NewBuffered(bytes.NewReader([]byte(" foo \nbar\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, " foo ", s)
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, "bar", s)

	r = read.NewBuffered(bytes.NewReader([]byte("\nfoo\n")))
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, "", s)
	s, err = r.Line()
	require.NoError(t, err)
	require.Equal(t, "foo", s)
}
