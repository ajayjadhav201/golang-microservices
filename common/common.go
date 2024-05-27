package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	Test bool = true
)

func NewUUID() string {
	return uuid.New().String()
}

// generateUniqueID generates a unique ID by combining milliseconds since epoch
// and a 6-digit random string.
func UniqueID() uint64 {
	// Get current time in milliseconds since the epoch
	millis := uint64(time.Now().UnixNano() / int64(time.Millisecond))

	// Generate a 6-digit random number
	rand.Seed(time.Now().UnixNano())       // Seed the random number generator
	randomNum := rand.Intn(90000) + 100000 // Generates a random integer between 100000 and 999999

	// Concatenate milliseconds and random number
	uniqueID, _ := strconv.ParseUint(fmt.Sprintf("%d%06d", millis, randomNum), 10, 64)

	return uniqueID
}
