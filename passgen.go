package passgen

import (
	"math/rand"
	"time"
)

func LowerCase(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	var result = make([]rune,n)

	// make unique seed for rand int in every program run
	rand.Seed(time.Now().UTC().UnixNano())

	for i := range result{
		result[i] = letters[rand.Intn(len(letters))] //rand.Intn for random int in 0-maks(n)
	}

	return string(result)
}