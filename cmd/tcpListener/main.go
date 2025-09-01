package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"tcp2http/internal/request"
	"tcp2http/internal/response"
	"tcp2http/internal/server"
)

const PORT = 8080

func equals(target, path string) bool {
	return target == path
}

func ConstructResponse(statusCode int, status, message string) (int, []byte) {
	body := []byte{}
	body = fmt.Appendf(body, `<html><head<title>%d %s</title</head<body<h1>%s</h1<p>%s</p</body</html>`, statusCode, status, status, message)
	length := len(body)
	return length, body
}

func main() {
	s, err := server.Connect(
		PORT,
		func(w *response.Writer, req *request.Request) {
			h := response.ConstructResponse("0")

			if equals(req.StartLine.Path, "/yourproblem") {
				message := "this is your problem buddy"
				w.WriteStatusLine(response.CODE_BAD_REQUEST)
				strNum, body := ConstructResponse(response.CODE_BAD_REQUEST, response.STATUS_BAD_REQUEST, message)
				n := strconv.Itoa(strNum)
				h.Replace(request.CONTENT_LENGTH_FORMAT, n)
				w.WriteHeaders(h)
				w.WriteBody(body)

			} else if equals(req.StartLine.Path, "/myproblem") {
				message := "my bad hehe~~"
				w.WriteStatusLine(response.CODE_INTERNAL_SERVER_ERROR)
				strNum, body := ConstructResponse(response.CODE_INTERNAL_SERVER_ERROR, response.STATUS_INTERNAL_SERVER_ERROR, message)
				n := strconv.Itoa(strNum)
				h.Replace(request.CONTENT_LENGTH_FORMAT, n)
				w.WriteHeaders(h)
				w.WriteBody(body)
			} else if strings.HasPrefix(req.StartLine.Path, "/httpbin/stream/") {
				target := req.StartLine.Path
				res, err := http.Get("https://httpbin.org/" + target[len("/httpbin/"):])
				if err != nil {
					message := "my bad hehe~~"
					w.WriteStatusLine(response.CODE_INTERNAL_SERVER_ERROR)
					strNum, body := ConstructResponse(response.CODE_INTERNAL_SERVER_ERROR, response.STATUS_INTERNAL_SERVER_ERROR, message)
					n := strconv.Itoa(strNum)
					h.Replace(request.CONTENT_LENGTH_FORMAT, n)
					w.WriteHeaders(h)
					w.WriteBody(body)
				} else {
					w.WriteStatusLine(response.CODE_OK)
					h.Delete(request.CONTENT_LENGTH_FORMAT)
					h.Set("transfer-encoding", "chunked")
					h.Replace("content-type", "text/plain")
					w.WriteHeaders(h)

					for {
						data := make([]byte, 32)
						n, err := res.Body.Read(data)
						if err != nil || n == 0 {
							break
						}
						body := []byte{}
						w.WriteBody(fmt.Appendf(body, "%x\r\n", n))
						w.WriteBody(data[:n])
						w.WriteBody(fmt.Append(body, "0\r\n\r\n"))
					}
					return
				}
			} else {
				w.WriteStatusLine(response.CODE_OK)
				message := "GG MAN!!"
				strNum, body := ConstructResponse(response.CODE_OK, response.STATUS_OK, message)
				n := strconv.Itoa(strNum)
				h.Replace(request.CONTENT_LENGTH_FORMAT, n)
				w.WriteHeaders(h)
				w.WriteBody(body)
			}
		})

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	defer s.Close()
	log.Println("Server started on port", PORT)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Server gracefully stopped")
}
