package user

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/KaiJi7/common/structs"
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
