package model

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestNewUserPassword_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系",
			input:    "password123",
			expected: "password123",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err := NewUserPassword(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			err = bcrypt.CompareHashAndPassword([]byte(actual), []byte(c.expected))
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}
		})
	}
}

func TestNewUserPassword_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "空文字の場合はエラー",
			input:    "",
			expected: "",
		},
		{
			title:    "7文字以下の場合はエラー",
			input:    "abcdefg",
			expected: "",
		},
		{
			title:    "31文字以上の場合はエラー",
			input: 	"abcdefghijklmnopqrstuvwxyz12345",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewUserPassword(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}
