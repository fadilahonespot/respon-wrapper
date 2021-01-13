package respon_wrapper

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"strings"
)

func EncryptedSha(inputs ...string) string {
	var userId string
	for _, input := range inputs {
		userId += input
	}
	h := sha1.New()
	h.Write([]byte(userId))
	var encrypted = h.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	return encryptedString
}

func HashPassword(input string) string {
	reader := strings.NewReader(input)
	hash := sha256.New()
	_, err := io.Copy(hash, reader)
	if err != nil {
		log.Println(err)
	}
	sum := hash.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", sum)
	return encryptedString
}

func RandomUid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x", b[0:4], b[6:8], b[4:6], b[10:])
	return uuid
}

func RandomPassword() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
	}
	uuid := fmt.Sprintf("%x%x", b[0:4], b[6:8])
	return uuid
}
