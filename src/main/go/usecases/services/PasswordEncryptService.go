package services

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomCharacter() string {
	character := Charset[seededRand.Intn(len(Charset))]
	return string(character)
}

func RandomSalt() string {
	result := ""
	for i:=0; i<6; i++ {
		result += RandomCharacter()
	}
	return result
}

func Encrypt(unEncryptedPassword string) string {
	hash := sha256.Sum256([]byte(unEncryptedPassword))
	return fmt.Sprintf("%x", hash[:])
}
