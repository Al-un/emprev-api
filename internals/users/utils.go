package users

import "crypto/sha512"

var (
	pwdSecretSalt string
)

func init() {
	pwdSecretSalt = "pouet"
}

// HashPassword hashes a password with the "pwdSecretSalt" which is appended to
// the password as a salt
func HashPassword(clearPassword string) string {
	h := sha512.New()
	h.Write([]byte(clearPassword))
	h.Write([]byte(pwdSecretSalt))
	hashedPassword := string(h.Sum(nil))

	return hashedPassword
	// return clearPassword
}
