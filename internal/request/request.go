package request

import (
	"bytes"
	"fmt"
	"httpOverTcp/internal/headers"
	"io"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

type State string

const INITIALIZED State = "Initialized"
const START_LINE_PARSED State = "Start Line Parsed"
const HEADERS_PARSED = "Headers Parsed"
const BODY_PARSED = "Body Parsed"

const BUFF_SIZE = 1024
const EMPTY_BODY_DEFAULT_CONTENT_LENGTH = 0
const CONTENT_LENGTH_FORMAT = "content-length"

var CRLF = []byte("\r\n")
var BAD_START_LINE = fmt.Errorf("Method, Path or HTTP Version Malformed or Missing")
var UNABLE_TO_READ_REQUEST = fmt.Errorf("Unable to read the entirety of the request")
var UNABLE_TO_PARSE_REQUEST = fmt.Errorf("Unable to parse the request")

type StartLine struct {
	Method      string
	HttpVersion string
	Path        string
}

type Request struct {
	StartLine StartLine
	Headers   *headers.Headers
	Body      string
	State     State
}

func newRequest() *Request {
	return &Request{
		Headers: headers.NewHeaders(),
		State:   INITIALIZED,
		Body:    "",
	}
}

func (r *Request) isBodyEmpty(key string) (bool, int, error) {
	lengthStr, err := r.Headers.Get(strings.ToLower(key))
	if err != nil {
		return true, EMPTY_BODY_DEFAULT_CONTENT_LENGTH, err
	}

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return true, EMPTY_BODY_DEFAULT_CONTENT_LENGTH, err
	}

	return false, length, nil
}

func (r *Request) parse(data []byte) (int, error) {

	start := 0
outer:

	for {
		rest := data[start:]
		slog.Info("rest", "rest", rest)
		switch r.State {
		case INITIALIZED:
			{
				l, redNum, err := RequestContentParser(rest)
				if err != nil {
					return 0, UNABLE_TO_PARSE_REQUEST
				}

				if redNum == 0 {
					break outer
				}

				r.State = START_LINE_PARSED
				r.StartLine = *l
				start += redNum
			}
		case START_LINE_PARSED:
			{
				redNum, done, err := r.Headers.ParseHeaders(rest)
				if err != nil {
					return 0, UNABLE_TO_PARSE_REQUEST
				}

				if redNum == 0 {
					break outer
				}

				if done {
					r.State = HEADERS_PARSED
					start += redNum + len(CRLF)
					break
				}

				start += redNum
			}
		case HEADERS_PARSED:
			{
				empty, l, err := r.isBodyEmpty(CONTENT_LENGTH_FORMAT)
				if err != nil {
					return 0, err
				}

				if empty || l == 0 {
					r.State = BODY_PARSED
					break outer
				}

				remaining := min(l-len(r.Body), len(rest))
				if remaining <= 0 {
					break outer
				}

				r.Body += string(rest[:remaining])
				start += remaining

				if len(r.Body) >= l {
					r.State = BODY_PARSED
					break outer
				}
			}
		case BODY_PARSED:
			{
				break outer
			}
		default:
			{
				return 0, UNABLE_TO_PARSE_REQUEST
			}
		}
	}

	return start, nil
}

func (s *StartLine) ValidateStartLine() bool {
	methodPattern, _ := regexp.Compile("[A-Za-z]{3,}")
	matched := methodPattern.FindString(s.Method)

	if s.HttpVersion != "HTTP/1.1" || (s.Path[0] != '/' && s.Path[0] != '*') || matched == "" {
		return false
	}

	return true
}

func RequestContentParser(content []byte) (*StartLine, int, error) {
	i := bytes.Index(content, CRLF)
	if i == -1 {
		return nil, 0, nil
	}

	startLine := content[:i]

	SlFields := bytes.Split(startLine, []byte(" "))
	if len(SlFields) != 3 {
		return nil, len(startLine), BAD_START_LINE
	}

	sl := &StartLine{
		Method:      strings.ToUpper(string(SlFields[0])),
		Path:        string(SlFields[1]),
		HttpVersion: string(SlFields[2]),
	}

	if !sl.ValidateStartLine() {
		return nil, len(startLine), BAD_START_LINE
	}

	return sl, len(startLine), nil
}

func RequestContentReader(reader io.Reader) (*Request, error) {
	req := newRequest()
	buff := make([]byte, BUFF_SIZE)
	buffStart := 0

	for req.State != BODY_PARSED {
		redNum, err := reader.Read(buff[buffStart:])

		if redNum == 0 {
			_, contentLen, _ := req.isBodyEmpty(CONTENT_LENGTH_FORMAT)

			if len(req.Body) < contentLen || len(req.Body) > contentLen {
				return nil, UNABLE_TO_PARSE_REQUEST
			}

			return req, nil
		}

		if err != nil {
			return nil, UNABLE_TO_READ_REQUEST
		}

		buffStart += redNum

		parsedNum, err := req.parse(buff[:buffStart])
		if err != nil {
			return nil, UNABLE_TO_PARSE_REQUEST
		}

		copy(buff, buff[parsedNum:buffStart])
		buffStart -= parsedNum
	}

	return req, nil
}
