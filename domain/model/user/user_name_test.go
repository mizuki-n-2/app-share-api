package model

import (
	"testing"
)

func TestNewUserName_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系",
			input:    "test name",
			expected: "test name",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err := NewUserName(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			if actual != UserName(c.expected) {
				t.Fatalf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestNewUserName_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "空文字の場合はエラー",
			input: "",
			expected: "",
		},
		{
			title:    "1文字以下の場合はエラー",
			input:    "a",
			expected: "",
		},
		{
			title:    "21文字以上の場合はエラー",
			input:    "abcdefghijklmnopqrstu",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewUserName(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}
