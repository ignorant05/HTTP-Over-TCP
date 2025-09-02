package tests

import (
	"httpOverTcp/internal/headers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeadersParsing(t *testing.T) {
	h := headers.NewHeaders()
	data := []byte("Host: localhost:8080\r\n\r\n")
	n, done, err := h.ParseHeaders(data)
	require.NoError(t, err)
	require.NotNil(t, h)
	res, err := h.Get("host")
	assert.Equal(t, "localhost:8080", res)
	require.Nil(t, err)
	assert.True(t, done)
	assert.Equal(t, 23, n)

	h = headers.NewHeaders()
	data = []byte("       Host : localhost:8080       \r\n\r\n")
	n, done, err = h.ParseHeaders(data)
	require.Error(t, err)
	assert.False(t, done)
	assert.Equal(t, 0, n)

	h = headers.NewHeaders()
	data = []byte("Host: localhost:8080\r\nHost: localhost:8080\r\nHost: localhost:8080\r\n\r\n")
	n, done, err = h.ParseHeaders(data)
	require.NoError(t, err)
	require.NotNil(t, h)
	res, err = h.Get("host")
	assert.Equal(t, "localhost:8080, localhost:8080, localhost:8080", res)
	require.Nil(t, err)
	assert.True(t, done)
	assert.Equal(t, 69, n)

	h = headers.NewHeaders()
	data = []byte("Host: localhost:8080\r\nContent-length: 122\r\nContent-type: text\r\n\r\n")
	n, done, err = h.ParseHeaders(data)
	require.NoError(t, err)
	require.NotNil(t, h)
	res, err = h.Get("host")
	assert.Equal(t, "localhost:8080", res)
	require.Nil(t, err)
	res, err = h.Get("content-length")
	assert.Equal(t, "122", res)
	require.Nil(t, err)
	res, err = h.Get("content-type")
	assert.Equal(t, "text", res)
	require.Nil(t, err)
	assert.True(t, done)
}
