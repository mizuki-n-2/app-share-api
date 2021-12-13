package test

import (
	"testing"
	"app-share-api/domain/model"
)

func TestNewUserNameSuccess(t *testing.T) {
	userName, err := model.NewUserName("test name")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if userName != "test name" {
		t.Fatal("failed test")
	}
}

func TestNewUserNameFailEmpty(t *testing.T) {
	_, err := model.NewUserName("")
	if err == nil {
		t.Fatal("failed test")
	}
}

// 名前が21文字以上の場合はエラー
func TestNewUserNameFailTooLong(t *testing.T) {
	_, err := model.NewUserName("abcdefghijklmnopqrstu")
	if err == nil {
		t.Fatal("failed test")
	}
}