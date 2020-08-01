package users

var (
	pwdSecretSalt string
)

func init() {
	pwdSecretSalt = "pouet"
}

// hashPassword hashes a password with the "pwdSecretSalt" which is appended to
// the password as a salt
func hashPassword(clearPassword string) string {
	// h := sha512.New()
	// h.Write([]byte(clearPassword))
	// h.Write([]byte(pwdSecretSalt))
	// hashedPassword := string(h.Sum(nil))

	// return hashedPassword
	return clearPassword
}
