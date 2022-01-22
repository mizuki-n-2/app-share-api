package model

import (
	"testing"
)

func TestNewUserEmail_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系1",
			input:    "test@example.com",
			expected: "test@example.com",
		},
		{
			title:    "正常系2",
			input:    "test@com",
			expected: "test@com",
		},
		{
			title:    "正常系3",
			input:    "test.sample@com",
			expected: "test.sample@com",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err := NewUserEmail(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			if actual != UserEmail(c.expected) {
				t.Fatalf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestNewUserEmail_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title: "空文字の場合はエラー",
			input: "",
			expected: "",
		},
		{
			title: "@がない場合はエラー",
			input: "test",
			expected: "",
		},
		{
			title: "@の後ろがない場合はエラー",
			input: "test@",
			expected: "",
		},
		{
			title: "@の前がない場合はエラー",
			input: "@example.com",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewUserEmail(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}

