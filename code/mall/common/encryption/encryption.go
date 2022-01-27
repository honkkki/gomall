package encryption

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string {
	key, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(key))
}
