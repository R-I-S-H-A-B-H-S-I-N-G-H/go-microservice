package encryption_util

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/hashing_util"
	"os"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func getSecureCookie() *securecookie.SecureCookie {
	var hashKey = []byte(os.Getenv("HASH_KEY"))
	var blockKey = []byte(os.Getenv("BLOCK_KEY"))
	if s == nil {
		s = securecookie.New(hashKey, blockKey)
	}
	return s
}

// EncryptData encrypts and encodes data using securecookie
func EncryptData(data interface{}) (string, error) {
	return getSecureCookie().Encode("data", data)
}

func DecryptData(encoded string) (string, error) {
	var decoded string
	err := getSecureCookie().Decode("data", encoded, &decoded)
	return decoded, err
}

func GenerateNewCookie() (string, error) {
	var err error
	cookie, err := hashing_util.GenerateRandomHash(8)
	if err != nil {
		return "", err
	}
	return EncryptData(cookie)
}
