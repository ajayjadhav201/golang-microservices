package common

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	Test bool = true
)

func EnvString(Key string, option string) string {
	val := os.Getenv(Key)
	// val := GetSecret(Key)
	if val != "" {
		return val
	}
	return option
}

func NewUUID() string {
	return uuid.New().String()
}

// generateUniqueID generates a unique ID by combining milliseconds since epoch
// and a 6-digit random string.
func UniqueID() int64 {
	// Generate a 6-digit random number
	randomNum := rand.Intn(9000) + 10000 // Generates a random integer between 100000 and 999999

	// Concatenate milliseconds and random number
	uniqueID, _ := strconv.ParseUint(fmt.Sprintf("%d%06d", time.Now().UnixMilli(), randomNum), 10, 64)

	return int64(uniqueID)
}

func Int64toa(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func Atoi(s string, defVal ...int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		if len(defVal) == 0 {
			return -1
		}
		return defVal[0]
	}
	return i
}
