package model

import (
	"testing"
)

func TestNewUserNameSuccess(t *testing.T) {
	SUCCESS_EXAMPLE_USER_NAME := "test name"

	userName, err := NewUserName(SUCCESS_EXAMPLE_USER_NAME)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if userName != UserName(SUCCESS_EXAMPLE_USER_NAME) {
		t.Fatal("failed test")
	}
}

func TestNewUserNameFailEmpty(t *testing.T) {
	_, err := NewUserName("")
	if err == nil {
		t.Fatal("failed test")
	}
}

// 名前が1文字以下の場合はエラー
func TestNewUserNameFailTooShort(t *testing.T) {
	FAIL_EXAMPLE_USER_NAME := "a"

	_, err := NewUserName(FAIL_EXAMPLE_USER_NAME)
	if err == nil {
		t.Fatal("failed test")
	}
}

// 名前が21文字以上の場合はエラー
func TestNewUserNameFailTooLong(t *testing.T) {
	FAIL_EXAMPLE_USER_NAME := "abcdefghijklmnopqrstu"

	_, err := NewUserName(FAIL_EXAMPLE_USER_NAME)
	if err == nil {
		t.Fatal("failed test")
	}
}