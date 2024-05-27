package common

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	InvalidReqBody = "Invalid request body."
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

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, map[string]string{"error": message})
}

func WriteInternalServerError(w http.ResponseWriter) {
	WriteError(w, http.StatusInternalServerError, "Internal server error")
}

func WriteServerNotAvailableError(w http.ResponseWriter) {
	WriteError(w, http.StatusInternalServerError, "Server not available.")
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
		} else {
			WriteError(w, http.StatusBadRequest, err.Error())
		}
	} else {
		WriteError(w, http.StatusBadRequest, err.Error())
	}

}

func WriteRequestBodyError(w http.ResponseWriter, err error) {
	//create custom error messages here
	if strings.Contains(err.Error(), "Key") {
		newerr := ""
		for i, err := range err.(validator.ValidationErrors) {
			if i != 0 {
				newerr = Sprintf("%s, ", newerr)
			}
			fieldName := err.Field()
			fieldTag := err.Tag()

			Println("ajaj fields are name: ", fieldName, "tag : ", fieldTag)
			switch fieldTag {
			case "email":
				Println("ajaj showing email error ", newerr)
				newerr = Sprintf("%s%s", newerr, "email is not valid")
			case "mobile":
				Println("ajaj showing mobile number error ", newerr)
				newerr = Sprintf("%s%s", newerr, "mobile number is not valid")
			case "required":
				newerr = Sprintf("%s%s is required", newerr, fieldName)
			case "gte":
				newerr = Sprintf("%s%s is too short", newerr, fieldName)
			}

		}
		if len(newerr) < 1 {
			WriteError(w, http.StatusBadRequest, InvalidReqBody)
		} else {
			WriteError(w, http.StatusBadRequest, newerr)
		}
		// return errors.New(newerr)
	} else if strings.Contains(err.Error(), "EOF") {
		WriteError(w, http.StatusBadRequest, InvalidReqBody)
	} else {
		WriteError(w, http.StatusBadRequest, InvalidReqBody)
	}
}
