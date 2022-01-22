package model

import (
	"testing"
)

func TestNewUserBio_Success(t *testing.T) {
	SUCCESS_EXAMPLE_USER_BIO := "こんにちは。よろしくお願いします！"

	userBio, err := NewUserBio(SUCCESS_EXAMPLE_USER_BIO)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if userBio != UserBio(SUCCESS_EXAMPLE_USER_BIO) {
		t.Fatal("failed test")
	}
}

// 自己紹介文が256文字以上の場合はエラー
func TestNewUserBio_FailTooLong(t *testing.T) {
	FAIL_EXAMPLE_USER_BIO := "こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！！"

	_, err := NewUserBio(FAIL_EXAMPLE_USER_BIO)
	if err == nil {
		t.Fatal("failed test")
	}
}