package headers

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var INVALID_HEADER_FORMAT = fmt.Errorf("Malformed header")
var HEADER_NOT_FOUND = fmt.Errorf("Header not found")
var CRLF = []byte("\r\n")
var DoubleCRLF = []byte("\r\n\r\n")

type Headers struct {
	Headers map[string]string
}

func NewHeaders() *Headers {
	return &Headers{
		Headers: map[string]string{},
	}
}

func (h *Headers) FormatHeaders(f func(key, val string)) {
	for key, val := range h.Headers {
		f(key, val)
	}
}

func (h *Headers) Get(fidldName string) (string, error) {
	if _, exits := h.Headers[fidldName]; !exits {
		return "", HEADER_NOT_FOUND
	}

	return h.Headers[fidldName], nil
}

func (h *Headers) Set(key, value string) {
	key = strings.ToLower(key)

	if orgValue, exits := h.Headers[key]; exits {
		h.Headers[key] = fmt.Sprintf("%s, %s", orgValue, value)
	} else {
		h.Headers[key] = value
	}
}

func (h *Headers) Replace(key, value string) {
	key = strings.ToLower(key)
	_, err := h.Get(key)
	if err != nil {
		h.Set(key, value)
	}

	key = strings.ToLower(key)
	h.Headers[key] = value
}

func (h *Headers) Delete(key string) error {
	key = strings.ToLower(key)

	_, err := h.Get(key)
	if err != nil {
		return HEADER_NOT_FOUND
	}

	delete(h.Headers, key)

	return nil
}

func isHeaderValid(header []byte) (bool, error) {
	validHeader, err := regexp.Compile(`^\s*([A-Za-z0-9_.$%#+*\|!~'-^]+):\s*(.+?)\s*(\r\n)?$`)
	if err != nil {
		return false, err
	}

	matched := validHeader.Find(header)
	if matched == nil {
		return false, INVALID_HEADER_FORMAT
	}

	return true, nil
}

func ParseFieldLine(header []byte) (string, string, error) {
	if valid, err := isHeaderValid(header); !valid {
		return "", "", err
	}

	pair := bytes.SplitN(header, []byte(":"), 2)
	if len(pair) != 2 {
		return "", "", INVALID_HEADER_FORMAT
	}

	key, value := bytes.TrimSpace(pair[0]), bytes.TrimSpace(pair[1])

	return string(key), string(value), nil
}

func (h *Headers) ParseHeaders(data []byte) (int, bool, error) {
	done, redIdx := false, 0

outer:
	for {
		eohIdx := bytes.Index(data[redIdx:], DoubleCRLF)
		if eohIdx == -1 {
			done = true
			break outer
		}

		currIdx := bytes.Index(data[redIdx:], CRLF)
		if currIdx == -1 {
			break outer
		}

		if currIdx == 0 {
			redIdx += len(CRLF)
			break
		}

		key, value, err := ParseFieldLine(data[redIdx : redIdx+currIdx])
		if err != nil {
			return 0, false, err
		}

		redIdx += currIdx + len(CRLF)
		h.Set(key, value)
	}

	return redIdx, done, nil
}
