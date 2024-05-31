package common

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const (
	UnmarshalFailed   = "Failed to parse request body"
	InternalServerErr = "Internal server error"
)

type Response struct {
	Message string `json:"message"`
}

func WriteJSON(c *gin.Context, status int, data any) {
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(status)
	// json.NewEncoder(w).Encode(data)
	c.JSON(status, data)
}

func ReadJSON(c *gin.Context, pointer any) error {
	err := json.NewDecoder(c.Request.Body).Decode(pointer)
	if err != nil {
		Println("ajaj readjson first error: ", err.Error())
		return err
	}
	err = validate.Struct(pointer)
	if err != nil {
		Println("ajaj readjson first error: ", err.Error())
		return err
	}
	return nil

	// if r.ContentLength <= 0 {
	// 	return errors.New("bad request")
	// }
	// b, e := io.ReadAll(r.Body)
	// if e != nil {
	// 	return errors.New("failed to read request body")
	// }
	// return json.Unmarshal(b, pointer)
}
