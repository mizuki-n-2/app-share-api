package model

import (
	"testing"
)

func TestNewPostTitle_Success(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "正常系",
			input:    "post title",
			expected: "post title",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			actual, err := NewPostTitle(c.input)
			if err != nil {
				t.Fatalf("failed test %#v", err)
			}

			if actual != PostTitle(c.expected) {
				t.Fatalf("expected: %v, actual: %v", c.expected, actual)
			}
		})
	}
}

func TestNewPostTitle_Fail(t *testing.T) {
	cases := []struct {
		title, input, expected string
	}{
		{
			title:    "空文字の場合はエラー",
			input: "",
			expected: "",
		},
		{
			title:    "101文字以上の場合はエラー",
			input:    "ここも途中ようやくその相当痛というのの中を合うでん。できるだけ今を反抗院はじっとある解釈ですたなどから嫌うてしまいないでは話籠っますだば、それほどには教えでたでしな。主意にやっつけた事はできるだけ前を同時にですうで。",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {
			_, err := NewPostTitle(c.input)
			if err == nil {
				t.Fatal("failed test")
			}
		})
	}
}
