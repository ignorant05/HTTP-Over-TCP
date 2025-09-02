package response

import (
	"fmt"
	"httpOverTcp/internal/headers"
	"io"
)

type StatusCode int

type Status string

type Writer struct {
	Writer io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Writer: w,
	}
}

// Status Codes
const (
	CODE_CONTINUE            = 100
	CODE_SWITCHING_PROTOCOLS = 101
	CODE_PROCESSING          = 102
	CODE_EARLY_HINTS         = 103

	CODE_OK                            = 200
	CODE_CREATED                       = 201
	CODE_ACCEPTED                      = 202
	CODE_NON_AUTHORITATIVE_INFORMATION = 203
	CODE_NO_CONTENT                    = 204
	CODE_RESET_CONTENT                 = 205
	CODE_PARTIAL_CONTENT               = 206
	CODE_MULTI_STATUS                  = 207
	CODE_ALREADY_REPORTED              = 208
	CODE_IM_USED                       = 226

	CODE_MULTIPLE_CHOICES   = 300
	CODE_MOVED_PERMANENTLY  = 301
	CODE_FOUND              = 302
	CODE_SEE_OTHER          = 303
	CODE_NOT_MODIFIED       = 304
	CODE_USE_PROXY          = 305
	CODE_TEMPORARY_REDIRECT = 307
	CODE_PERMANENT_REDIRECT = 308

	CODE_BAD_REQUEST                     = 400
	CODE_UNAUTHORIZED                    = 401
	CODE_PAYMENT_REQUIRED                = 402
	CODE_FORBIDDEN                       = 403
	CODE_NOT_FOUND                       = 404
	CODE_METHOD_NOT_ALLOWED              = 405
	CODE_NOT_ACCEPTABLE                  = 406
	CODE_PROXY_AUTHENTICATION_REQUIRED   = 407
	CODE_REQUEST_TIMEOUT                 = 408
	CODE_CONFLICT                        = 409
	CODE_GONE                            = 410
	CODE_LENGTH_REQUIRED                 = 411
	CODE_PRECONDITION_FAILED             = 412
	CODE_PAYLOAD_TOO_LARGE               = 413
	CODE_URI_TOO_LONG                    = 414
	CODE_UNSUPPORTED_MEDIA_TYPE          = 415
	CODE_RANGE_NOT_SATISFIABLE           = 416
	CODE_EXPECTATION_FAILED              = 417
	CODE_IM_A_TEAPOT                     = 418
	CODE_MISDIRECTED_REQUEST             = 421
	CODE_UNPROCESSABLE_ENTITY            = 422
	CODE_LOCKED                          = 423
	CODE_FAILED_DEPENDENCY               = 424
	CODE_TOO_EARLY                       = 425
	CODE_UPGRADE_REQUIRED                = 426
	CODE_PRECONDITION_REQUIRED           = 428
	CODE_TOO_MANY_REQUESTS               = 429
	CODE_REQUEST_HEADER_FIELDS_TOO_LARGE = 431
	CODE_UNAVAILABLE_FOR_LEGAL_REASONS   = 451

	CODE_INTERNAL_SERVER_ERROR           = 500
	CODE_NOT_IMPLEMENTED                 = 501
	CODE_BAD_GATEWAY                     = 502
	CODE_SERVICE_UNAVAILABLE             = 503
	CODE_GATEWAY_TIMEOUT                 = 504
	CODE_HTTP_VERSION_NOT_SUPPORTED      = 505
	CODE_VARIANT_ALSO_NEGOTIATES         = 506
	CODE_INSUFFICIENT_STORAGE            = 507
	CODE_LOOP_DETECTED                   = 508
	CODE_NOT_EXTENDED                    = 510
	CODE_NETWORK_AUTHENTICATION_REQUIRED = 511
)

// Status Strings
const (
	STATUS_CONTINUE            = "Continue"
	STATUS_SWITCHING_PROTOCOLS = "Switching Protocols"
	STATUS_PROCESSING          = "Processing"
	STATUS_EARLY_HINTS         = "Early Hints"

	STATUS_OK                            = "OK"
	STATUS_CREATED                       = "Created"
	STATUS_ACCEPTED                      = "Accepted"
	STATUS_NON_AUTHORITATIVE_INFORMATION = "Non-Authoritative Information"
	STATUS_NO_CONTENT                    = "No Content"
	STATUS_RESET_CONTENT                 = "Reset Content"
	STATUS_PARTIAL_CONTENT               = "Partial Content"
	STATUS_MULTI_STATUS                  = "Multi-Status"
	STATUS_ALREADY_REPORTED              = "Already Reported"
	STATUS_IM_USED                       = "IM Used"

	STATUS_MULTIPLE_CHOICES   = "Multiple Choices"
	STATUS_MOVED_PERMANENTLY  = "Moved Permanently"
	STATUS_FOUND              = "Found"
	STATUS_SEE_OTHER          = "See Other"
	STATUS_NOT_MODIFIED       = "Not Modified"
	STATUS_USE_PROXY          = "Use Proxy"
	STATUS_TEMPORARY_REDIRECT = "Temporary Redirect"
	STATUS_PERMANENT_REDIRECT = "Permanent Redirect"

	STATUS_BAD_REQUEST                     = "Bad Request"
	STATUS_UNAUTHORIZED                    = "Unauthorized"
	STATUS_PAYMENT_REQUIRED                = "Payment Required"
	STATUS_FORBIDDEN                       = "Forbidden"
	STATUS_NOT_FOUND                       = "Not Found"
	STATUS_METHOD_NOT_ALLOWED              = "Method Not Allowed"
	STATUS_NOT_ACCEPTABLE                  = "Not Acceptable"
	STATUS_PROXY_AUTHENTICATION_REQUIRED   = "Proxy Authentication Required"
	STATUS_REQUEST_TIMEOUT                 = "Request Timeout"
	STATUS_CONFLICT                        = "Conflict"
	STATUS_GONE                            = "Gone"
	STATUS_LENGTH_REQUIRED                 = "Length Required"
	STATUS_PRECONDITION_FAILED             = "Precondition Failed"
	STATUS_PAYLOAD_TOO_LARGE               = "Payload Too Large"
	STATUS_URI_TOO_LONG                    = "URI Too Long"
	STATUS_UNSUPPORTED_MEDIA_TYPE          = "Unsupported Media Type"
	STATUS_RANGE_NOT_SATISFIABLE           = "Range Not Satisfiable"
	STATUS_EXPECTATION_FAILED              = "Expectation Failed"
	STATUS_IM_A_TEAPOT                     = "I'm a teapot"
	STATUS_MISDIRECTED_REQUEST             = "Misdirected Request"
	STATUS_UNPROCESSABLE_ENTITY            = "Unprocessable Entity"
	STATUS_LOCKED                          = "Locked"
	STATUS_FAILED_DEPENDENCY               = "Failed Dependency"
	STATUS_TOO_EARLY                       = "Too Early"
	STATUS_UPGRADE_REQUIRED                = "Upgrade Required"
	STATUS_PRECONDITION_REQUIRED           = "Precondition Required"
	STATUS_TOO_MANY_REQUESTS               = "Too Many Requests"
	STATUS_REQUEST_HEADER_FIELDS_TOO_LARGE = "Request Header Fields Too Large"
	STATUS_UNAVAILABLE_FOR_LEGAL_REASONS   = "Unavailable For Legal Reasons"

	STATUS_INTERNAL_SERVER_ERROR           = "Internal Server Error"
	STATUS_NOT_IMPLEMENTED                 = "Not Implemented"
	STATUS_BAD_GATEWAY                     = "Bad Gateway"
	STATUS_SERVICE_UNAVAILABLE             = "Service Unavailable"
	STATUS_GATEWAY_TIMEOUT                 = "Gateway Timeout"
	STATUS_HTTP_VERSION_NOT_SUPPORTED      = "HTTP Version Not Supported"
	STATUS_VARIANT_ALSO_NEGOTIATES         = "Variant Also Negotiates"
	STATUS_INSUFFICIENT_STORAGE            = "Insufficient Storage"
	STATUS_LOOP_DETECTED                   = "Loop Detected"
	STATUS_NOT_EXTENDED                    = "Not Extended"
	STATUS_NETWORK_AUTHENTICATION_REQUIRED = "Network Authentication Required"
)

const (
	CONTENT_LENGTH_HEADER_FORMAT = "content-length"
	CONTENT_TYPE_HEADER_FORMAT   = "content-type"
	CONNECTION_HEADER_FORMAT     = "connection"
)

var (
	NOT_VALID                   = fmt.Errorf("Not a valid status code")
	UNABLE_TO_WRITE_TO_RESPONSE = fmt.Errorf("Unable to write to response body")
)

func (w *Writer) WriteStatusLine(statusCode StatusCode) error {
	statusLine := []byte{}

	switch statusCode {
	case CODE_CONTINUE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_CONTINUE, STATUS_CONTINUE)
		}
	case CODE_SWITCHING_PROTOCOLS:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_SWITCHING_PROTOCOLS, STATUS_SWITCHING_PROTOCOLS)
		}
	case CODE_PROCESSING:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PROCESSING, STATUS_PROCESSING)
		}
	case CODE_EARLY_HINTS:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_EARLY_HINTS, STATUS_EARLY_HINTS)
		}
	case CODE_OK:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_OK, STATUS_OK)
		}
	case CODE_CREATED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_CREATED, STATUS_CREATED)
		}
	case CODE_ACCEPTED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_ACCEPTED, STATUS_ACCEPTED)
		}
	case CODE_NON_AUTHORITATIVE_INFORMATION:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NON_AUTHORITATIVE_INFORMATION, STATUS_NON_AUTHORITATIVE_INFORMATION)
		}
	case CODE_NO_CONTENT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NO_CONTENT, STATUS_NO_CONTENT)
		}
	case CODE_RESET_CONTENT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_RESET_CONTENT, STATUS_RESET_CONTENT)
		}
	case CODE_PARTIAL_CONTENT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PARTIAL_CONTENT, STATUS_PARTIAL_CONTENT)
		}
	case CODE_MULTI_STATUS:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_MULTI_STATUS, STATUS_MULTI_STATUS)
		}
	case CODE_ALREADY_REPORTED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_ALREADY_REPORTED, STATUS_ALREADY_REPORTED)
		}
	case CODE_IM_USED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_IM_USED, STATUS_IM_USED)
		}
	case CODE_MULTIPLE_CHOICES:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_MULTIPLE_CHOICES, STATUS_MULTIPLE_CHOICES)
		}
	case CODE_MOVED_PERMANENTLY:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_MOVED_PERMANENTLY, STATUS_MOVED_PERMANENTLY)
		}
	case CODE_FOUND:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_FOUND, STATUS_FOUND)
		}
	case CODE_SEE_OTHER:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_SEE_OTHER, STATUS_SEE_OTHER)
		}
	case CODE_NOT_MODIFIED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NOT_MODIFIED, STATUS_NOT_MODIFIED)
		}
	case CODE_USE_PROXY:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_USE_PROXY, STATUS_USE_PROXY)
		}
	case CODE_TEMPORARY_REDIRECT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_TEMPORARY_REDIRECT, STATUS_TEMPORARY_REDIRECT)
		}
	case CODE_PERMANENT_REDIRECT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PERMANENT_REDIRECT, STATUS_PERMANENT_REDIRECT)
		}
	case CODE_BAD_REQUEST:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_BAD_REQUEST, STATUS_BAD_REQUEST)
		}
	case CODE_UNAUTHORIZED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_UNAUTHORIZED, STATUS_UNAUTHORIZED)
		}
	case CODE_PAYMENT_REQUIRED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PAYMENT_REQUIRED, STATUS_PAYMENT_REQUIRED)
		}
	case CODE_FORBIDDEN:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_FORBIDDEN, STATUS_FORBIDDEN)
		}
	case CODE_NOT_FOUND:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NOT_FOUND, STATUS_NOT_FOUND)
		}
	case CODE_METHOD_NOT_ALLOWED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_METHOD_NOT_ALLOWED, STATUS_METHOD_NOT_ALLOWED)
		}
	case CODE_NOT_ACCEPTABLE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NOT_ACCEPTABLE, STATUS_NOT_ACCEPTABLE)
		}
	case CODE_PROXY_AUTHENTICATION_REQUIRED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PROXY_AUTHENTICATION_REQUIRED, STATUS_PROXY_AUTHENTICATION_REQUIRED)
		}
	case CODE_REQUEST_TIMEOUT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_REQUEST_TIMEOUT, STATUS_REQUEST_TIMEOUT)
		}
	case CODE_CONFLICT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_CONFLICT, STATUS_CONFLICT)
		}
	case CODE_GONE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_GONE, STATUS_GONE)
		}
	case CODE_LENGTH_REQUIRED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_LENGTH_REQUIRED, STATUS_LENGTH_REQUIRED)
		}
	case CODE_PRECONDITION_FAILED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PRECONDITION_FAILED, STATUS_PRECONDITION_FAILED)
		}
	case CODE_PAYLOAD_TOO_LARGE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PAYLOAD_TOO_LARGE, STATUS_PAYLOAD_TOO_LARGE)
		}
	case CODE_URI_TOO_LONG:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_URI_TOO_LONG, STATUS_URI_TOO_LONG)
		}
	case CODE_UNSUPPORTED_MEDIA_TYPE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_UNSUPPORTED_MEDIA_TYPE, STATUS_UNSUPPORTED_MEDIA_TYPE)
		}
	case CODE_RANGE_NOT_SATISFIABLE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_RANGE_NOT_SATISFIABLE, STATUS_RANGE_NOT_SATISFIABLE)
		}
	case CODE_EXPECTATION_FAILED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_EXPECTATION_FAILED, STATUS_EXPECTATION_FAILED)
		}
	case CODE_IM_A_TEAPOT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_IM_A_TEAPOT, STATUS_IM_A_TEAPOT)
		}
	case CODE_MISDIRECTED_REQUEST:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_MISDIRECTED_REQUEST, STATUS_MISDIRECTED_REQUEST)
		}
	case CODE_UNPROCESSABLE_ENTITY:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_UNPROCESSABLE_ENTITY, STATUS_UNPROCESSABLE_ENTITY)
		}
	case CODE_LOCKED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_LOCKED, STATUS_LOCKED)
		}
	case CODE_FAILED_DEPENDENCY:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_FAILED_DEPENDENCY, STATUS_FAILED_DEPENDENCY)
		}
	case CODE_TOO_EARLY:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_TOO_EARLY, STATUS_TOO_EARLY)
		}
	case CODE_UPGRADE_REQUIRED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_UPGRADE_REQUIRED, STATUS_UPGRADE_REQUIRED)
		}
	case CODE_PRECONDITION_REQUIRED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_PRECONDITION_REQUIRED, STATUS_PRECONDITION_REQUIRED)
		}
	case CODE_TOO_MANY_REQUESTS:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_TOO_MANY_REQUESTS, STATUS_TOO_MANY_REQUESTS)
		}
	case CODE_REQUEST_HEADER_FIELDS_TOO_LARGE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_REQUEST_HEADER_FIELDS_TOO_LARGE, STATUS_REQUEST_HEADER_FIELDS_TOO_LARGE)
		}
	case CODE_UNAVAILABLE_FOR_LEGAL_REASONS:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_UNAVAILABLE_FOR_LEGAL_REASONS, STATUS_UNAVAILABLE_FOR_LEGAL_REASONS)
		}
	case CODE_INTERNAL_SERVER_ERROR:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_INTERNAL_SERVER_ERROR, STATUS_INTERNAL_SERVER_ERROR)
		}
	case CODE_NOT_IMPLEMENTED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NOT_IMPLEMENTED, STATUS_NOT_IMPLEMENTED)
		}
	case CODE_BAD_GATEWAY:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_BAD_GATEWAY, STATUS_BAD_GATEWAY)
		}
	case CODE_SERVICE_UNAVAILABLE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_SERVICE_UNAVAILABLE, STATUS_SERVICE_UNAVAILABLE)
		}
	case CODE_GATEWAY_TIMEOUT:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_GATEWAY_TIMEOUT, STATUS_GATEWAY_TIMEOUT)
		}
	case CODE_HTTP_VERSION_NOT_SUPPORTED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_HTTP_VERSION_NOT_SUPPORTED, STATUS_HTTP_VERSION_NOT_SUPPORTED)
		}
	case CODE_VARIANT_ALSO_NEGOTIATES:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_VARIANT_ALSO_NEGOTIATES, STATUS_VARIANT_ALSO_NEGOTIATES)
		}
	case CODE_INSUFFICIENT_STORAGE:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_INSUFFICIENT_STORAGE, STATUS_INSUFFICIENT_STORAGE)
		}
	case CODE_LOOP_DETECTED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_LOOP_DETECTED, STATUS_LOOP_DETECTED)
		}
	case CODE_NOT_EXTENDED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NOT_EXTENDED, STATUS_NOT_EXTENDED)
		}
	case CODE_NETWORK_AUTHENTICATION_REQUIRED:
		{
			statusLine = fmt.Appendf(statusLine, "HTTP/1.1 %d %s", CODE_NETWORK_AUTHENTICATION_REQUIRED, STATUS_NETWORK_AUTHENTICATION_REQUIRED)
		}
	default:
		{
			return NOT_VALID
		}
	}

	statusLine = fmt.Append(statusLine, "\r\n")

	_, err := w.Writer.Write(statusLine)
	if err != nil {
		return UNABLE_TO_WRITE_TO_RESPONSE
	}
	return nil
}

func ConstructResponse(contentLen string) *headers.Headers {
	h := headers.NewHeaders()
	h.Set(CONTENT_LENGTH_HEADER_FORMAT, contentLen)
	h.Set(CONTENT_TYPE_HEADER_FORMAT, "text/plain")
	h.Set(CONNECTION_HEADER_FORMAT, "close")

	return h
}

func (w *Writer) WriteHeaders(h *headers.Headers) error {
	bytesHeaders := []byte{}

	h.FormatHeaders(func(key, val string) {
		bytesHeaders = fmt.Appendf(bytesHeaders, "%s: %s\r\n", key, val)
	})

	bytesHeaders = fmt.Append(bytesHeaders, "\r\n")
	_, err := w.Writer.Write(bytesHeaders)
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) WriteBody(body []byte) (int, error) {
	n, err := w.Writer.Write(body)
	return n, err
}
