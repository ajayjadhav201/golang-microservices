package common

import (
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Fatal(err error, s ...string) {
	if err != nil {
		if len(s) == 0 {
			log.Fatal(err)
		} else {
			log.Fatal(s)
		}
	}
}

func Panic(err error, s ...string) {
	if err != nil {
		if len(s) == 0 {
			panic(err)
		} else {
			panic(s)
		}
	}
}

func WriteGrpcError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	// Check if the error is a gRPC status error
	if st, ok := status.FromError(err); ok {
		// Extract the gRPC status code
		grpcCode := st.Code()
		if grpcCode == codes.Unavailable {
			WriteServerNotAvailableError(w)
		}
	} else {
		WriteInternalServerError(w)
	}

}
