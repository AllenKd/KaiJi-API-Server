package user

import (
	"KaiJi-Admin/internal/pkg/db/collection"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestSignIn(t *testing.T) {
	testUser := structs.User{
		Name: "testUser",
		Password: "password",
	}

	h := sha256.New()
	h.Write([]byte(testUser.Password))
	hashedPassword := hex.EncodeToString(h.Sum(nil))
	_ = hashedPassword
}
