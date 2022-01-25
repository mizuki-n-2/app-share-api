package model

import (
	"testing"
)

func TestNewPostContent_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系",
			input:    "このアプリはSNSです。",
			expected: "このアプリはSNSです。",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err := NewPostContent(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			if actual != PostContent(c.expected) {
				t.Fatalf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestNewPostContent_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "256文字以上の場合はエラー",
			input:    "このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。このアプリはSNSです。",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewPostContent(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}
