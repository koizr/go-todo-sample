package jwt

import (
	"github.com/koizr/go-todo-sample/auth/domain"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	//goland:noinspection SpellCheckingInspection
	secret := "HKb0vrVNMFCWoaAPH9X7n3GUCBM3+EzmDEEL9y9tGSFnWMKI3x0ohg=="
	user := domain.User{ID: "55f3bb74-57ca-40ed-b566-5a98e2efb0eb"}
	now := time.Date(2020, 10, 1, 12, 40, 30, 50, time.UTC)

	signedString, err := GenerateToken(secret, &user, &now)
	if err != nil {
		t.Fatalf("failed to generate token")
	}

	//goland:noinspection SpellCheckingInspection
	expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDE1NTYzMzAsImlhdCI6MTYwMTU1NjAzMCwic3ViIjoiNTVmM2JiNzQtNTdjYS00MGVkLWI1NjYtNWE5OGUyZWZiMGViIn0.haVNKjAcUShzZUUtUuZsk1z4RTV19hTHMi_RGa_RNAU"
	if signedString != expected {
		t.Errorf("token mismatch. signedString=%s", signedString)
	}
}
