package str

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const alphabet = "./ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

var bcEncoding = base64.NewEncoding(alphabet)

func GenerateUUID() string {
	return uuid.NewString()
}

func GenerateSalt() string {
	unencodedSalt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, unencodedSalt)
	if err != nil {
		return ""
	}

	return string(base64Encode(unencodedSalt))
}

func HashStr(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)

	return string(bytes), err
}

func CompareHash(hashed, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str))
	return err == nil
}

func RandStr(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func base64Encode(src []byte) []byte {
	n := bcEncoding.EncodedLen(len(src))
	dst := make([]byte, n)
	bcEncoding.Encode(dst, src)
	for dst[n-1] == '=' {
		n--
	}
	return dst[:n]
}
