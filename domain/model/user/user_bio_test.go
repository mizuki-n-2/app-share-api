package model

import (
	"testing"
)

func TestNewUserBio_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系",
			input:    "こんにちは。よろしくお願いします！",
			expected: "こんにちは。よろしくお願いします！",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err := NewUserBio(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			if actual != UserBio(c.expected) {
				t.Fatalf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestNewUserBio_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "256文字以上の場合はエラー",
			input:    "こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！こんにちは。よろしくお願いします！！",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewUserBio(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}
