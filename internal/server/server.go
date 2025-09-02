package server

import (
	"fmt"
	"httpOverTcp/internal/request"
	"httpOverTcp/internal/response"
	"io"
	"net"
)

type HandlerError struct {
	StatusCode response.StatusCode
	Err        string
}
type Handler func(w *response.Writer, req *request.Request)

type State string

const OFF = "No Connection"
const ON = "Connected"

type Server struct {
	State   State
	Handler Handler
}

func NewServer(handler Handler) *Server {
	return &Server{
		State:   ON,
		Handler: handler,
	}
}

func Run(s *Server, conn io.ReadWriteCloser) {
	defer conn.Close()
	h := response.ConstructResponse("0")

	writer := response.NewWriter(conn)
	r, err := request.RequestContentReader(conn)
	if err != nil {
		msg := []byte("sorry my bad, it's my problem this time")
		writer.WriteStatusLine(response.CODE_INTERNAL_SERVER_ERROR)
		h.Replace(request.CONTENT_LENGTH_FORMAT, fmt.Sprintf("%d", len(msg)))
		writer.WriteHeaders(h)
		writer.WriteBody(msg)
		return
	}

	s.Handler(writer, r)
}

func Connect(PORT uint16, handler Handler) (*Server, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		return nil, err
	}

	s := NewServer(handler)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil || s.State == OFF {
				return
			}

			go Run(s, conn)
		}
	}()

	return s, nil
}

func (s *Server) Close() error {
	s.State = OFF
	return nil
}
