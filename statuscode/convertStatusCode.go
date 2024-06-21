package statuscode

import (
	"google.golang.org/grpc/codes"
)

// |-----------------------------------|---------------------------|
// | gRPC Status Code                  | HTTP Status Code          |
// |-----------------------------------|---------------------------|
// |`OK` (Code: 0)               	   | 200 OK                    |
// |`CANCELLED` (Code: 1)        	   | 499 Client Closed Request |
// |`UNKNOWN` (Code: 2)                | 500 Internal Server Error |
// |`INVALID_ARGUMENT` (Code: 3) 	   | 400 Bad Request           |
// |`DEADLINE_EXCEEDED` (Code: 4)	   | 504 Gateway Timeout       |
// |`NOT_FOUND` (Code: 5)       	   | 404 Not Found             |
// |`ALREADY_EXISTS` (Code: 6)         | 409 Conflict              |
// |`PERMISSION_DENIED` (Code: 7)      | 403 Forbidden             |
// |`RESOURCE_EXHAUSTED`(Code:8)       | 429 Too Many Requests     |
// |`FAILED_PRECONDITION`(Code:9)      | 400 Bad Request           |
// |`ABORTED` (Code: 10)               | 409 Conflict              |
// |`OUT_OF_RANGE` (Code: 11)          | 400 Bad Request           |
// |`UNIMPLEMENTED` (Code: 12)         | 501 Not Implemented       |
// |`INTERNAL` (Code: 13)              | 500 Internal Server Error |
// |`UNAVAILABLE` (Code: 14)           | 503 Service Unavailable   |
// |`DATA_LOSS` (Code: 15)             | 500 Internal Server Error |
// |`UNAUTHENTICATED` (Code: 16)       | 401 Unauthorized          |
// |-----------------------------------|---------------------------|

// ConvertGrpcToHTTP converts gRPC status code (of type codes.code) to HTTP status code
//
// Note: Maximum gRPC status code is 16 when this function is written. For other status code, it will return 500 (http.StatusInternalServerError)
func ConvertGrpcToHTTP(sCode codes.Code) int {
	switch sCode {
	case codes.OK:	//Code: 0
		return 200 //http.StatusOK
	case codes.Canceled:	//Code: 1
		return 499
	case codes.Unknown:	//Code: 2
		return 500 //http.StatusInternalServerError
	case codes.InvalidArgument:	//Code: 3
		return 400 //http.StatusBadRequest
	case codes.DeadlineExceeded:	//Code: 4
		return 504 //http.StatusGatewayTimeout
	case codes.NotFound:	//Code: 5
		return 404 //http.StatusNotFound
	case codes.AlreadyExists:	//Code: 6
		return 409 //http.StatusConflict
	case codes.PermissionDenied:	//Code: 7
		return 403 //http.StatusForbidden
	case codes.ResourceExhausted:	//Code: 8	
		return 429 //http.StatusTooManyRequests
	case codes.FailedPrecondition:	//Code: 9	
		return 400 //http.StatusBadRequest
	case codes.Aborted:	//Code: 10
		return 409 //http.StatusConflict
	case codes.OutOfRange:	//Code: 11
		return 400 //http.StatusBadRequest
	case codes.Unimplemented:	//Code: 12
		return 501 //http.StatusNotImplemented
	case codes.Internal:	//Code: 13
		return 500 //http.StatusInternalServerError
	case codes.Unavailable:	//Code: 14
		return 503 //http.StatusServiceUnavailable
	case codes.DataLoss:	//Code: 15
		return 500 //http.StatusInternalServerError
	case codes.Unauthenticated:	//Code: 16
		return 401 //http.StatusUnauthorized
	default:
		return 500 //http.StatusInternalServerError
	}
}