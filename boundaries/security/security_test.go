package security

import (
	"testing"
)

func TestSecurityImpl(t *testing.T) {
	s := NewSecurityImpl()
	password := "my_secret_password"

	// Test Encoding
	encoded, err := s.EncodePassword(password)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if encoded == "" || encoded == password {
		t.Fatalf("Expected encoded password to be different from original and not empty")
	}

	// Test Correct Comparison
	match, err := s.ComparePassword(password, encoded)
	if err != nil {
		t.Fatalf("Expected no error on comparison, got %v", err)
	}
	if !match {
		t.Fatal("Expected passwords to match")
	}

	// Test Incorrect Comparison
	match, err = s.ComparePassword("wrong_password", encoded)
	if err != nil {
		t.Fatalf("Expected no error on comparison with wrong password, got %v", err)
	}
	if match {
		t.Fatal("Expected passwords NOT to match")
	}
}
