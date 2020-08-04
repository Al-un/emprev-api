package users

import (
	"crypto/sha512"
	"os"
)

var (
	pwdSecretSalt string
)

func init() {
	pwdSecretSalt = "super-secret-password-salt"
	if secretVar := os.Getenv("SECRET_PWD_SALT"); secretVar != "" {
		pwdSecretSalt = secretVar
	}
}

// HashPassword hashes a password with the "pwdSecretSalt" which is appended to
// the password as a salt. Hash is done with SHA-512.
func HashPassword(clearPassword string) string {
	h := sha512.New()
	h.Write([]byte(clearPassword))
	h.Write([]byte(pwdSecretSalt))
	hashedPassword := string(h.Sum(nil))

	return hashedPassword
}
