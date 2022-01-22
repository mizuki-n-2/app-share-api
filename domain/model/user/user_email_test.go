package model

import (
	"testing"
)

func TestNewUserEmail_Success(t *testing.T) {
	SUCCESS_EXAMPLE_USER_EMAIL := "test@example.com"

	userEmail, err := NewUserEmail(SUCCESS_EXAMPLE_USER_EMAIL)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if userEmail != UserEmail(SUCCESS_EXAMPLE_USER_EMAIL) {
		t.Fatal("failed test")
	}
}

func TestNewUserEmail_FailEmpty(t *testing.T) {
	_, err := NewUserEmail("")
	if err == nil {
		t.Fatal("failed test")
	}
}

// メールアドレスの形式が不正の場合はエラー(@がない)
func TestNewUserName_FailFormat1(t *testing.T) {
	FAIL_EXAMPLE_USER_EMAIL := "test"

	_, err := NewUserEmail(FAIL_EXAMPLE_USER_EMAIL)
	if err == nil {
		t.Fatal("failed test")
	}
}

// メールアドレスの形式が不正の場合はエラー(@の後ろがない)
func TestNewUserName_FailFormat2(t *testing.T) {
	FAIL_EXAMPLE_USER_EMAIL := "test@"

	_, err := NewUserEmail(FAIL_EXAMPLE_USER_EMAIL)
	if err == nil {
		t.Fatal("failed test")
	}
}

// メールアドレスの形式が不正の場合はエラー(@の前がない)
func TestNewUserName_FailFormat3(t *testing.T) {
	FAIL_EXAMPLE_USER_EMAIL := "@com"

	_, err := NewUserEmail(FAIL_EXAMPLE_USER_EMAIL)
	if err == nil {
		t.Fatal("failed test")
	}
}
