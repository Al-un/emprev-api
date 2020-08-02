package users

import (
	"fmt"
	"testing"
)

func TestPouet(t *testing.T) {
	fmt.Printf(HashPassword(dbSuperAdminUserName))
}
