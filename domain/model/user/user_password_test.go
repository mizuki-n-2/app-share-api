package model

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNewUserPasswordSuccess(t *testing.T) {
	SUCCESS_EXAMPLE_USER_PASSWORD := "password123"

	userPassword, err := NewUserPassword(SUCCESS_EXAMPLE_USER_PASSWORD)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(SUCCESS_EXAMPLE_USER_PASSWORD))
	if err != nil {
		t.Fatal("failed test")
	}
}

func TestNewUserPasswordFailEmpty(t *testing.T) {
	_, err := NewUserPassword("")
	if err == nil {
		t.Fatal("failed test")
	}
}

// パスワードが7文字以下の場合はエラー
func TestNewUserPasswordFailTooShort(t *testing.T) {
	FAIL_EXAMPLE_USER_PASSWORD := "abcdefg"

	_, err := NewUserPassword(FAIL_EXAMPLE_USER_PASSWORD)
	if err == nil {
		t.Fatal("failed test")
	}
}

// パスワードが31文字以上の場合はエラー
func TestNewUserPasswordFailTooLong(t *testing.T) {
	FAIL_EXAMPLE_USER_PASSWORD := "abcdefghijklmnopqrstuvwxyz12345"

	_, err := NewUserPassword(FAIL_EXAMPLE_USER_PASSWORD)
	if err == nil {
		t.Fatal("failed test")
	}
}