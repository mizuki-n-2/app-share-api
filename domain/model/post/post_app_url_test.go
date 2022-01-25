package model

import (
	"testing"
)

func TestNewPostAppURL_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系1",
			input:    "https://example.com",
			expected: "https://example.com",
		},
		{
			title:    "正常系2",
			input:    "http://example.com",
			expected: "http://example.com",
		},
		{
			title: "正常系3",
			input: "https://test",
			expected: "https://test",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err :=NewPostAppURL(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			if actual != PostAppURL(c.expected) {
				t.Fatalf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestNewPostAppURL_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title: "http://またはhttps://がない場合はエラー",
			input: "example.com",
			expected: "",
		},
		{
			title: "http://またはhttps://の後ろがない場合はエラー",
			input: "http://",
			expected: "",
		},
		{
			title: "http://またはhttps://ではない場合はエラー",
			input: "ftp://sample.com",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewPostAppURL(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}

