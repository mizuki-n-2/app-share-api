package model

import (
	"testing"
)

func TestNewUserName_Success(t *testing.T) {
	SUCCESS_EXAMPLE_USER_NAME := "test name"

	userName, err := NewUserName(SUCCESS_EXAMPLE_USER_NAME)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if userName != UserName(SUCCESS_EXAMPLE_USER_NAME) {
		t.Fatal("failed test")
	}
}

func TestNewUserName_FailEmpty(t *testing.T) {
	_, err := NewUserName("")
	if err == nil {
		t.Fatal("failed test")
	}
}

// 名前が1文字以下の場合はエラー
func TestNewUserName_FailTooShort(t *testing.T) {
	FAIL_EXAMPLE_USER_NAME := "a"

	_, err := NewUserName(FAIL_EXAMPLE_USER_NAME)
	if err == nil {
		t.Fatal("failed test")
	}
}

// 名前が21文字以上の場合はエラー
func TestNewUserName_FailTooLong(t *testing.T) {
	FAIL_EXAMPLE_USER_NAME := "abcdefghijklmnopqrstu"

	_, err := NewUserName(FAIL_EXAMPLE_USER_NAME)
	if err == nil {
		t.Fatal("failed test")
	}
}