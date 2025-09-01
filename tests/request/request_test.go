package tests

import (
	"io"
	"tcp2http/internal/request"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type chunkReader struct {
	data            string
	numBytesPerRead int
	pos             int
}

func (cr *chunkReader) Read(p []byte) (n int, err error) {
	if cr.pos >= len(cr.data) {
		return 0, io.EOF
	}
	endIndex := cr.pos + cr.numBytesPerRead
	endIndex = min(len(cr.data), endIndex)
	n = copy(p, cr.data[cr.pos:endIndex])
	cr.pos += n

	return n, nil
}

func TestRequestLineParse(t *testing.T) {
	// works
	reader := &chunkReader{
		data:            "GET / HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 3,
	}
	r, err := request.RequestContentReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.StartLine.Method)
	assert.Equal(t, "/", r.StartLine.Path)
	assert.Equal(t, "HTTP/1.1", r.StartLine.HttpVersion)

	// don't work
	reader = &chunkReader{
		data:            "GET /coffee HTTP/1.1\r\nHost: localhost:8080\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 1,
	}
	r, err = request.RequestContentReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "GET", r.StartLine.Method)
	assert.Equal(t, "/coffee", r.StartLine.Path)
	assert.Equal(t, "HTTP/1.1", r.StartLine.HttpVersion)
}

func TestRequestHeadersParse(t *testing.T) {
	// works
	reader := &chunkReader{
		data:            "GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n",
		numBytesPerRead: 3,
	}
	r, err := request.RequestContentReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	r1, err := r.Headers.Get("host")
	assert.Equal(t, "localhost:42069", r1)
	require.NotNil(t, r1)
	r2, err := r.Headers.Get("user-agent")
	assert.Equal(t, "curl/7.81.0", r2)
	require.NotNil(t, r2)
	r3, err := r.Headers.Get("accept")
	assert.Equal(t, "*/*", r3)
	require.NotNil(t, r3)

	// don't work
	reader = &chunkReader{
		data:            "GET / HTTP/1.1\r\nHost localhost:42069\r\n\r\n",
		numBytesPerRead: 3,
	}
	r, err = request.RequestContentReader(reader)
	require.Error(t, err)
}
func TestReqFinal(t *testing.T) {
	// works
	reader := &chunkReader{
		data: "POST /submit HTTP/1.1\r\n" +
			"Host: localhost:8080\r\n" +
			"Content-Length: 13\r\n" +
			"\r\n" +
			"hello world!\n",
		numBytesPerRead: 3,
	}
	r, err := request.RequestContentReader(reader)
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "hello world!\n", string(r.Body))

	// don't work
	reader = &chunkReader{
		data: "POST /submit HTTP/1.1\r\n" +
			"Host: localhost:8080\r\n" +
			"Content-Length: 20\r\n" +
			"\r\n" +
			"partial content",
		numBytesPerRead: 3,
	}
	r, err = request.RequestContentReader(reader)
	require.Error(t, err)

}
