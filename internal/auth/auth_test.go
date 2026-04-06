package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	tokenString, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT() error: %v", err)
	}

	gotID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateJWT() error: %v", err)
	}
	if gotID != userID {
		t.Errorf("ValidateJWT() = %v, want %v", gotID, userID)
	}
}

func TestValidateJWT_ExpiredToken(t *testing.T) {
	userID := uuid.New()
	secret := "test-secret"

	tokenString, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT() error: %v", err)
	}

	_, err = ValidateJWT(tokenString, secret)
	if err == nil {
		t.Fatal("ValidateJWT() expected error for expired token, got nil")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	userID := uuid.New()

	tokenString, err := MakeJWT(userID, "correct-secret", time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT() error: %v", err)
	}

	_, err = ValidateJWT(tokenString, "wrong-secret")
	if err == nil {
		t.Fatal("ValidateJWT() expected error for wrong secret, got nil")
	}
}

func TestGetBearerToken(t *testing.T) {
	tests := []struct {
		name      string
		header    string
		wantToken string
		wantErr   bool
	}{
		{"valid token", "Bearer my-token-123", "my-token-123", false},
		{"empty header", "", "", true},
		{"missing prefix", "my-token-123", "", true},
		{"wrong prefix", "Basic my-token-123", "", true},
		{"extra whitespace", "Bearer   my-token-123  ", "my-token-123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}
			got, err := GetBearerToken(headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.wantToken {
				t.Errorf("GetBearerToken() = %q, want %q", got, tt.wantToken)
			}
		})
	}
}
